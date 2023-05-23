#!/bin/bash
set -e

# TODO remove configuration from here
KMS_KEY=c42rf64v82v478nnn1vt
ENDPOINT=api.il.nebius.cloud:443
AUTH_KEY_FILE=../../auth_key_il.json

START_DIR=$(pwd)
trap 'cd $START_DIR' EXIT

SCRIPT_PATH=$(dirname "${BASH_SOURCE[0]}")
cd "${WORK_DIR:-$SCRIPT_PATH}"
cd vault

echo "Building vault"
make bootstrap
make dev

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
  endpoint = "${ENDPOINT:-api.cloud.yandex.net:443}"
  service_account_key_file = "${AUTH_KEY_FILE:-auth_key.json}"
}
EOF

if [[ -d "./storage" ]]; then
  echo "Cleaning storage"
  rm -rf ./storage/*
fi

./bin/vault server -config vault.hcl -log-level=error -log-file=vault.log &
PID=$!
trap 'kill $PID' EXIT
# waiting for server to start
sleep 5

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
echo "SUCCESS"
