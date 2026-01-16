#!/bin/bash
set -e

SCRIPT_PATH=$(dirname "${BASH_SOURCE[0]}")
. $SCRIPT_PATH/../common.sh
. $SCRIPT_PATH/test_local.cfg

PID=0
cleanup() {
    yc --profile $PROFILE kms symmetric-key delete --folder-id $FOLDER_ID vault-test-local
    yc --profile $PROFILE iam service-account delete --folder-id $FOLDER_ID --name "test-local-vault-sa"
    kill $PID
    exit
}

trap "cleanup" EXIT;

init
init_vault
cd vault

echo "Testing $YCKMS_VERSION branch"
git checkout $YCKMS_VERSION

echo "Create service account"
yc --profile $PROFILE iam service-account create --folder-id $FOLDER_ID --name "test-local-vault-sa"

echo "Create KMS key"
KMS_KEY=$(yc --profile $PROFILE kms symmetric-key create --folder-id $FOLDER_ID --name vault-test-local --default-algorithm aes-256 | yq '.id')

echo "Create auth key for service account"
yc --profile $PROFILE iam key create --folder-id $FOLDER_ID --service-account-name test-local-vault-sa --output ./auth_key.json

echo "Give roles for service account"
yc --profile $PROFILE resource-manager folder add-access-binding $FOLDER_ID \
   --role kms.keys.encrypterDecrypter \
   --service-account-name test-local-vault-sa \
   --folder-id $FOLDER_ID

echo "Building vault"
make bootstrap
make dev

echo $KMS_ENDPOINT

cat >vault.hcl <<EOF
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
  service_account_key_file = "auth_key.json"
}
EOF

if [[ -d "./storage" ]]; then
  echo "Cleaning storage"
  rm -rf ./storage/*
fi

./bin/vault server -config vault.hcl -log-level=error -log-file=vault.log &
PID=$!
# waiting for server to start
sleep 10

export VAULT_ADDR=http://127.0.0.1:8200

echo "Initializing vault"
export VAULT_TOKEN=$(./bin/vault operator init | grep "Initial Root Token:" | cut -f 4 -d " ")
echo "Vault key: $VAULT_TOKEN"

echo "Enabling key value storage"
./bin/vault secrets enable -path=secret/ kv

KEY=foo
VAL=bar
echo "Writing key-value"
./bin/vault write secret/my-secret $KEY=$VAL

echo "Reading key-value"
ACTUAL_VAL=$(./bin/vault read -field "$KEY" secret/my-secret)

if [[ "$VAL" != "$ACTUAL_VAL" ]]; then
  echo >&2 "Invalid key '$KEY' value! Expected $VAL, but was $ACTUAL_VAL"
  exit 1
fi
echo "Local test passed successfully"
