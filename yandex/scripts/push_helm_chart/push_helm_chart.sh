#!/bin/bash
set -e

SCRIPT_PATH=$(dirname "${BASH_SOURCE[0]}")
  . $SCRIPT_PATH/helm_chart.cfg

if [[ -n $WORK_DIR ]]; then
  mkdir -p $WORK_DIR
  cd $WORK_DIR
fi

IAM_TOKEN=$(ycp --profile="$PROFILE" iam create-token)
DOMAIN="cr.yandex"

if [[ "$PROFILE" == "israel" ]]; then
  DOMAIN="cr.cloudil.com"
fi

if [[ "$PROFILE" == "preprod" ]]; then
  DOMAIN="cr.cloud-preprod.yandex.net"
fi

docker login --username iam --password $IAM_TOKEN $DOMAIN

git clone https://github.com/hashicorp/vault-helm.git ./vault-helm
git clone ssh://git@bb.yandexcloud.net/cloud/mk8s-marketplace-helm.git ./mk8s-marketplace-helm

cd vault-helm

TAGS=$(git tag --sort="-version:refname")
echo $TAGS

VERSION=$(echo "$BASE_VERSION" | cut -c 2-)
CUT_VERSION="${VERSION%.*}"
ACTUAL_TAG=''

for TAG in $TAGS
do
  git checkout tags/$TAG values.yaml
  ACTUAL_VERSION=$(yq '.server.image.tag' values.yaml)
  ACTUAL_VERSION="${ACTUAL_VERSION%.*}"
  echo "Actual version = $ACTUAL_VERSION and version = $CUT_VERSION"
  if [ $CUT_VERSION = $ACTUAL_VERSION ]
  then
    ACTUAL_TAG=$TAG
    echo "Found actual version $ACTUAL_VERSION"
    break
  fi
done

git checkout tags/$ACTUAL_TAG

sudo rm -r ./.git

CSI_PROVIDER_ACTUAL_TAG=$(yq '.csi.image.tag' values.yaml)
K8S_ACTUAL_TAG=$(yq '.injector.image.tag' values.yaml)

cd ../

docker pull hashicorp/vault-csi-provider:$CSI_PROVIDER_ACTUAL_TAG --platform amd64
docker pull hashicorp/vault-k8s:$K8S_ACTUAL_TAG --platform amd64

CSI_PROVIDER_IMAGE_ID=$(docker images hashicorp/vault-csi-provider:$CSI_PROVIDER_ACTUAL_TAG -q)
K8S_IMAGE_ID=$(docker images hashicorp/vault-k8s:$K8S_ACTUAL_TAG -q)

docker tag $K8S_IMAGE_ID $DOMAIN/$REGISTRY_ID/vault/vault-k8s:$K8S_ACTUAL_TAG
docker tag $K8S_IMAGE_ID $DOMAIN/$REGISTRY_ID/vault/vault-k8s:latest

docker tag $CSI_PROVIDER_IMAGE_ID $DOMAIN/$REGISTRY_ID/vault/vault-csi-provider:$CSI_PROVIDER_ACTUAL_TAG
docker tag $CSI_PROVIDER_IMAGE_ID $DOMAIN/$REGISTRY_ID/vault/vault-csi-provider:latest

docker push $DOMAIN/$REGISTRY_ID/vault/vault-k8s:$K8S_ACTUAL_TAG
docker push $DOMAIN/$REGISTRY_ID/vault/vault-k8s:latest

docker push $DOMAIN/$REGISTRY_ID/vault/vault-csi-provider:$CSI_PROVIDER_ACTUAL_TAG
docker push $DOMAIN/$REGISTRY_ID/vault/vault-csi-provider:latest

rm -r ./mk8s-marketplace-helm/products/hashicorp-vault/chart
cp -r ./vault-helm ./mk8s-marketplace-helm/products/hashicorp-vault/chart
rm -r ./mk8s-marketplace-helm/products/hashicorp-vault/chart/test

CHART_VERSION="$(yq '.version' ./vault-helm/Chart.yaml)-1"
replacement=$CHART_VERSION yq -i '.version = strenv(replacement)' ./mk8s-marketplace-helm/products/hashicorp-vault/chart/Chart.yaml

cat >./mk8s-marketplace-helm/products/hashicorp-vault/chart/templates/kms-creds-secret.yaml <<EOF
apiVersion: v1
kind: Secret
metadata:
  name: kms-creds
  namespace: {{ .Release.Namespace | quote }}
  labels:
    chart: "{{ .Chart.Name }}-{{ .Chart.Version | replace "+" "_" }}"
type: Opaque
data:
  "credentials.json": {{ .Values.yandexKmsAuthJson | b64enc }}
EOF

/bin/bash $SCRIPT_PATH/update_values.sh

cd mk8s-marketplace-helm/products/hashicorp-vault/chart
helm package .
helm push ./vault-$(echo "$ACTUAL_TAG" | cut -c 2-)-1.tgz oci://$DOMAIN/$REGISTRY_ID/vault/chart