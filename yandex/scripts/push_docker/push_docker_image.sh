#!/bin/bash
set -e

SCRIPT_PATH=$(dirname "${BASH_SOURCE[0]}")
  . $SCRIPT_PATH/../common.sh
  . $SCRIPT_PATH/push_docker.cfg

init
init_vault
cd vault

if ! git checkout $YCKMS_VERSION; then
    echo >&2 "Cannot find branch $YCKMS_VERSION !"
    exit 1
else
    echo "Checkout to branch '$YCKMS_VERSION' succeeded"
fi

cd yandex
cd docker

IAM_TOKEN=$(ycp --profile="$PROFILE" iam create-token)
DOMAIN="cr.yandex"

if [[ "$PROFILE" == "israel" ]]; then
  DOMAIN="cr.cloudil.com"
fi

if [[ "$PROFILE" == "preprod" ]]; then
  DOMAIN="cr.cloud-preprod.yandex.net"
fi

docker login --username iam --password $IAM_TOKEN $DOMAIN
docker build --platform linux/amd64 -t $DOMAIN/$REGISTRY_ID/vault:$BASE_VERSION-yckms -t $DOMAIN/$REGISTRY_ID/vault .
docker push $DOMAIN/$REGISTRY_ID/vault:$BASE_VERSION-yckms
docker logout