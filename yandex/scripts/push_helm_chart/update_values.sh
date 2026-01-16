#!/bin/bash

SCRIPT_PATH=$(dirname "${BASH_SOURCE[0]}")
  . $SCRIPT_PATH/release_sample.cfg

CSI_PROVIDER_ACTUAL_TAG=$(yq '.csi.image.tag' ./vault-helm/values.yaml)
K8S_ACTUAL_TAG=$(yq '.injector.image.tag' ./vault-helm/values.yaml)

yq '.' ./mk8s-marketplace-helm/products/hashicorp-vault/chart/values.yaml > values.yaml.new

DOMAIN="cr.yandex"
if [[ "$PROFILE" == "israel" ]]; then
  DOMAIN="cr.cloudil.com"
fi
if [[ "$PROFILE" == "preprod" ]]; then
  DOMAIN="cr.cloud-preprod.yandex.net"
fi
VAULT_VERSION="$(echo "$BASE_VERSION" | cut -c 2-)-yckms"

replacement=$VAULT_VERSION yq -i '(.. | select(key == "repository" and . == "hashicorp/vault") | parent).tag |= strenv(replacement)' values.yaml.new
replacement="$DOMAIN/$REGISTRY_ID/vault/vault" yq -i '(.. | select(key == "repository" and . == "hashicorp/vault") | parent).repository |= strenv(replacement)' values.yaml.new
replacement=$CSI_PROVIDER_ACTUAL_TAG yq -i '(.. | select(key == "repository" and . == "hashicorp/vault-csi-provider") | parent).tag |= strenv(replacement)' values.yaml.new
replacement="$DOMAIN/$REGISTRY_ID/vault/vault-csi-provider" yq -i '(.. | select(key == "repository" and . == "hashicorp/vault-csi-provider") | parent).repository |= strenv(replacement)' values.yaml.new
replacement=$K8S_ACTUAL_TAG yq -i '(.. | select(key == "repository" and . == "hashicorp/vault-k8s") | parent).tag |= strenv(replacement)' values.yaml.new
replacement="$DOMAIN/$REGISTRY_ID/vault/vault-k8s" yq -i '(.. | select(key == "repository" and . == "hashicorp/vault-k8s") | parent).repository |= strenv(replacement)' values.yaml.new
yq -i '(.. | select(key == "extraVolumes")) |= {"type": "secret", "name": "kms-creds"} ' values.yaml.new

yq '.' ./mk8s-marketplace-helm/products/hashicorp-vault/chart/values.yaml > values.yaml.noblanks
diff -B values.yaml.noblanks values.yaml.new > ./patch.file
patch ./mk8s-marketplace-helm/products/hashicorp-vault/chart/values.yaml ./patch.file

#patch works bad, so it's needed to patch twice
yq '.' ./mk8s-marketplace-helm/products/hashicorp-vault/chart/values.yaml > values.yaml.new

read -rd '' replacement << EOF

seal "yandexcloudkms" {
    kms_key_id = "{{ .Values.yandexKmsKeyId}}"
    service_account_key_file = "/vault/userconfig/kms-creds/credentials.json"
}

EOF
replacement=$replacement yq -i '(.. | select(key == "standalone")).config |= . + strenv(replacement)' values.yaml.new

yq '.' ./mk8s-marketplace-helm/products/hashicorp-vault/chart/values.yaml > values.yaml.noblanks
diff -B values.yaml.noblanks values.yaml.new > ./patch.file
patch ./mk8s-marketplace-helm/products/hashicorp-vault/chart/values.yaml ./patch.file

echo 'yandexKmsAuthJson: ""' >> ./mk8s-marketplace-helm/products/hashicorp-vault/chart/values.yaml
echo 'yandexKmsKeyId: ""' >> ./mk8s-marketplace-helm/products/hashicorp-vault/chart/values.yaml

perl -pe 'chomp if eof' ./mk8s-marketplace-helm/products/hashicorp-vault/chart/values.yaml >tmp.file
mv tmp.file ./mk8s-marketplace-helm/products/hashicorp-vault/chart/values.yaml

rm patch.file
rm values.yaml.new
rm values.yaml.noblanks
rm ./mk8s-marketplace-helm/products/hashicorp-vault/chart/values.yaml.orig