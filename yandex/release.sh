#!/bin/bash
set -e

SCRIPT_PATH=$(dirname "${BASH_SOURCE[0]}")
. $SCRIPT_PATH/common.sh
. $SCRIPT_PATH/release.cfg

init
init_vault
get_kms_wrapper_version

cd vault

echo "Pushing synchronized main"
git push origin main

echo "Updating yckms"
git checkout yckms
git rebase origin main
git push origin yckms

if git checkout $YCKMS_VERSION; then
  echo "Branch '$YCKMS_VERSION' already exists"
else
  echo "Creating branch '$YCKMS_VERSION'"
  git checkout -b $YCKMS_VERSION $BASE_VERSION
fi

PATCH_LAST_COMMIT_MSG="YCKMS patch"
while read -r line < <(git log $BASE_VERSION..$YCKMS_VERSION --oneline --reverse --pretty=format:"%B"); do
  if [[ $line == "$PATCH_LAST_COMMIT_MSG" ]]; then
    HAS_YCKMS_PATCH=true
    break
  fi
done

if [[ "$HAS_YCKMS_PATCH" != true ]]; then
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
  YCKMS_WRAPPER_VERSION="$KMS_WRAPPER_VERSION-$WRAPPER_SUFFIX"
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

  echo "Committing"
  git commit -m "$PATCH_LAST_COMMIT_MSG"
else
  echo "Patch is already applied"
fi

git push -f origin $YCKMS_VERSION
