#!/bin/bash
set -e

SCRIPT_PATH=$(dirname "${BASH_SOURCE[0]}")
. $SCRIPT_PATH/common.sh
. $SCRIPT_PATH/release_sample.cfg

init
init_vault
init_vault_kms_wrapper
get_kms_wrapper_version

BRANCH=release/${KMS_WRAPPER_VERSION}+"yckms"
TAG=${KMS_WRAPPER_VERSION}-"yckms"

cd vault-kms-wrapper

if git ls-remote --exit-code origin $BRANCH; then
  echo >&2 "Remote branch '$BRANCH' already exists!"
  echo >&2 "Update suffix 'yckms' to release new version"
  exit 1
fi

if git ls-remote --exit-code origin $TAG; then
  echo >&2 "Remote tag '$TAG' already exists!"
  echo >&2 "Update suffix 'yckms' to release new version"
  exit 1
fi

echo "Getting github.com/hashicorp/go-kms-wrapping/v2@$KMS_WRAPPER_VERSION"
go get github.com/hashicorp/go-kms-wrapping/v2@$KMS_WRAPPER_VERSION
go mod tidy
echo "Testing"
go test
git add go.mod go.sum

if ! git diff --cached --quiet --exit-code; then
  echo "Committing"
  git commit -m "Version updated $VERSION"
else
  echo "Nothing to commit"
fi

if git show-ref --quiet $BRANCH; then
  git branch -D $BRANCH
fi

echo "Creating branch $BRANCH"
git checkout -b $BRANCH

go mod vendor
git add vendor

if ! git diff --cached --quiet --exit-code; then
  echo "Committing vendor"
  git commit -m "Vendor"
else
  echo "Nothing to commit"
fi

git tag $TAG -f

git push origin $BRANCH
git push origin $TAG

git checkout main
git push origin main
