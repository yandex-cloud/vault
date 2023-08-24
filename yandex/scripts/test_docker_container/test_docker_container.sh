#!/bin/bash
set -e

SCRIPT_PATH=$(dirname "${BASH_SOURCE[0]}")
  . $SCRIPT_PATH/../common.sh
  . $SCRIPT_PATH/test_docker.cfg

cleanup() {
    yc --profile $PROFILE kms symmetric-key delete --folder-id $FOLDER_ID vault-test
    yc --profile $PROFILE iam service-account delete --folder-id $FOLDER_ID --name "test-vault-sa"
    yc --profile $PROFILE compute instance delete --folder-id $FOLDER_ID --name test-vault
    yc --profile $PROFILE vpc subnet delete --folder-id $FOLDER_ID --name test-vault-subnet
    yc --profile $PROFILE vpc network delete --folder-id $FOLDER_ID --name test-vault-network
    exit
}

trap "cleanup" EXIT;

echo "Create service account"
yc --profile $PROFILE iam service-account create --folder-id $FOLDER_ID --name "test-vault-sa"

echo "Create KMS key"
KMS_KEY=$(yc --profile $PROFILE kms symmetric-key create --folder-id $FOLDER_ID --name vault-test --default-algorithm aes-256 | yq '.id')

echo "Create auth key for service account"
yc iam key create --folder-id $FOLDER_ID --service-account-name test-vault-sa --output ./vol/key.json

echo "Give roles for service account"
yc --profile $PROFILE resource-manager folder add-access-binding $FOLDER_ID \
   --role kms.keys.encrypterDecrypter \
   --service-account-name test-vault-sa \
   --folder-id $FOLDER_ID

yc --profile $PROFILE resource-manager folder add-access-binding $FOLDER_ID \
    --role container-registry.images.puller \
    --service-account-name test-vault-sa \
    --folder-id $FOLDER_ID

echo "Create network"
ycp --profile $PROFILE vpc network create --name test-vault-network --folder-id $FOLDER_ID

echo "Create subnet"
yc --profile $PROFILE vpc subnet create \
  --name test-vault-subnet \
  --description "test vault subnet" \
  --network-name test-vault-network \
  --zone il1-a \
  --folder-id $FOLDER_ID \
  --range 192.168.0.0/24

echo "Create instance"
yc --profile $PROFILE compute instance create test-vault \
  --zone il1-a \
  --platform standard-v3 \
  --create-boot-disk  image-folder-id=standard-images,image-family=ubuntu-2004-lts \
  --cores 2 \
  --core-fraction 100 \
  --memory 2 \
  --ssh-key $PUBLIC_SSH \
  --folder-id $FOLDER_ID \
  --service-account-name test-vault-sa \
  --public-ip

echo "Get machine's ip"
MACHINE_IP=$(yc --profile $PROFILE compute instance get --name test-vault --folder-id $FOLDER_ID | yq '.network_interfaces[0].primary_v4_address.one_to_one_nat.address')

echo "Wait until status is RUNNING"
#status=''
#while [[ $status != "RUNNING" ]]
#do
#status=$(yc compute instance get --name test-vault --folder-id $FOLDER_ID | yq '.status')
#sleep 3
#done

sleep 20

DOMAIN="cr.yandex"

if [[ "$PROFILE" == "israel" ]]; then
  DOMAIN="cr.cloudil.com"
fi

if [[ "$PROFILE" == "preprod" ]]; then
  DOMAIN="cr.cloud-preprod.yandex.net"
fi

cat >./vol/vault.hcl <<EOF
# See https://www.vaultproject.io/docs/configuration for more details about configuration options

ui = true

storage "file" {
  path = "storage"
}

# HTTP listener (insecure)
listener "tcp" {
  address = "127.0.0.1:8200"
  tls_disable = 1
}

## Auto Unseal via Yandex Key Management Service (see https://cloud.yandex.ru/docs/kms/solutions/vault-secret for more details)
seal "yandexcloudkms" {
  kms_key_id = "$KMS_KEY"
  endpoint = "${KMS_ENDPOINT:-api.cloud.yandex.net:443}"
  service_account_key_file = "/vault/config/key.json"
}
EOF

echo "Send test files"
#send auth key for service account
scp -i $PRIVATE_SSH ./vol/key.json ./vol/vault.hcl ./vol/startup.sh yc-user@$MACHINE_IP:/home/yc-user

ssh -i $PRIVATE_SSH -o ServerAliveInterval=60 yc-user@$MACHINE_IP "export DOMAIN=$DOMAIN; export DOCKER_ENDPOINT=$DOCKER_ENDPOINT; export REGISTRY_ID=$REGISTRY_ID; export BASE_VERSION=$BASE_VERSION; bash -s" < script_inside_ssh.sh
