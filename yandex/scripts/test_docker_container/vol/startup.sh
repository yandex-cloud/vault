#!/bin/bash
set -e

vault server -config vault.hcl &
PID=$!
trap 'kill $PID' EXIT

echo "Initializing vault"
token=''
while [ ! $token ]
do
sleep 3
token=$(vault operator init | grep "Initial Root Token:" | cut -f 4 -d " ")
done
echo "Vault key: $VAULT_TOKEN"

export VAULT_TOKEN=$token

echo "Enabling key value storage"
vault secrets enable -path=secret/ kv

KEY=foo
VAL=bar
echo "Writing key-value"
vault write secret/my-secret $KEY=$VAL

echo "Reading key-value"
ACTUAL_VAL=$(vault read -field "$KEY" secret/my-secret)

if [[ "$VAL" != "$ACTUAL_VAL" ]]; then
  echo >&2 "Invalid key '$KEY' value! Expected $VAL, but was $ACTUAL_VAL"
  exit 1
fi
echo "Local test passed successfully"