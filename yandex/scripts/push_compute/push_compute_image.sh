#!/bin/bash
set -e

SCRIPT_PATH=$(dirname "${BASH_SOURCE[0]}")
. $SCRIPT_PATH/../common.sh
. $SCRIPT_PATH/push_compute.cfg

init
init_vault
cd vault

if ! git checkout $YCKMS_VERSION; then
    echo >&2 "Cannot find branch $YCKMS_VERSION !"
    exit 1
else
    echo "Checkout to branch '$YCKMS_VERSION' succeeded"
fi

TOKEN=$(ycp --profile $PROFILE iam create-token)

cd yandex/compute

if [[ $REPLACE_DEFAULT_ENDPOINT == 1 ]]; then
  echo "$(jq --arg ENDPOINT "$ENDPOINT" --arg ZONE "$ZONE" --arg PLATFORM_ID "$PLATFORM_ID" \
   '.builders[0] += {"endpoint":$ENDPOINT, "zone":$ZONE, "platform_id":$PLATFORM_ID}' vault.packer.json)" > vault.packer.json
fi

echo "$(jq --arg SOURCE_IMAGE_FAMILY "ubuntu-2004-lts" '.builders[0].source_image_family = $SOURCE_IMAGE_FAMILY' vault.packer.json)" > vault.packer.json

FOLDER_ID="$FOLDER_ID" TOKEN="$TOKEN" $SCRIPT_PATH/packer build vault.packer.json