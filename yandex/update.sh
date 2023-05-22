#!/bin/bash
set -e

# TODO remove configuration from here
BASE_VERSION="v1.13.2"
YCKMS_VERSION="$BASE_VERSION+yckms"

START_DIR=$(pwd)
trap 'cd $START_DIR' EXIT

SCRIPT_PATH=$(dirname "${BASH_SOURCE[0]}")
cd "$SCRIPT_PATH"

if [[ ! -d "vault" ]]; then
  echo "Cloning vault"
  git clone git@github.com:yandex-cloud/vault.git
  cd vault
  git remote add upstream git@github.com:hashicorp/vault.git
#git push origin main
else
  echo "Vault already cloned"
  cd vault
fi

git reset --hard

echo "Synchronizing with upstream"
git checkout main
git pull upstream main

echo "Fetching tags"
git fetch upstream --tags

echo "Updating yckms"
git checkout yckms
git rebase origin main
#git push origin yckms

if git checkout $YCKMS_VERSION; then
  echo "Branch '$YCKMS_VERSION' already exists"
else
  echo "Creating branch '$YCKMS_VERSION'"
  git checkout -b $YCKMS_VERSION $BASE_VERSION
fi

PATCH_LAST_COMMIT_MSG="YCKMS patch"
LAST_COMMIT_MSG=$(git log -n 1 --format=%B)
if [[ $LAST_COMMIT_MSG != $PATCH_LAST_COMMIT_MSG ]]; then
  echo "Applying patch from yckms branch"
  git cherry-pick --no-commit $(git log main..yckms -1000 --oneline --reverse --pretty=format:"%h" | paste -sd' ' -)
  # cherry-pick is more stable then merge-base
  #git diff $(git merge-base --fork-point main yckms) yckms | git apply
  sed -i '' 's/.*VersionMetadata.*=.*""/VersionMetadata = "yckms"/' version/version_base.go
  go fmt version/version_base.go
  git add version/version_base.go
  sed -i '' "s/ARG BASE_VAULT_VERSION=.*/ARG BASE_VAULT_VERSION=$BASE_VERSION/" yandex/docker/Dockerfile
  sed -i '' "s/BASE_VAULT_VERSION=.*/BASE_VAULT_VERSION=$BASE_VERSION/" yandex/compute/install.sh
  git add yandex/docker/Dockerfile yandex/compute/install.sh

  echo "Adding github.com/yandex-cloud/vault-kms-wrapper/v2 dependency"
  KMS_WRAPPER_VERSION=$(go list -m github.com/hashicorp/go-kms-wrapping/v2 | cut -f 2 -d " ")
  YCKMS_WRAPPER_VERSION="$KMS_WRAPPER_VERSION-yckms"
  YCKMS_WRAPPER=github.com/yandex-cloud/vault-kms-wrapper/v2@"$YCKMS_WRAPPER_VERSION"

  if ! go list -m "$YCKMS_WRAPPER"; then
    echo >&2 "Cannot find $YCKMS_WRAPPER, possible release required!"
    exit 1
  fi
  go mod edit -require="$YCKMS_WRAPPER"
  go mod tidy
  git add go.mod go.sum

  echo "Vendoring"
  go mod vendor
  git add vendor

  git commit -m $PATCH_LAST_COMMIT_MSG
else
  echo "Patch is already applied"
fi

#git push origin $YCKMS_VERSION

echo "Building vault"
make bootstrap
make dev
