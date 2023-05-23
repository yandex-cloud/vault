#!/bin/bash
set -e

VERSION="v2.0.7"
SUFFIX="yckms"
BRANCH=release/${VERSION}+${SUFFIX}
TAG=${VERSION}-${SUFFIX}

START_DIR=$(pwd)
trap 'cd $START_DIR' EXIT

SCRIPT_PATH=$(dirname "${BASH_SOURCE[0]}")
cd "${WORK_DIR:-$SCRIPT_PATH}"

if [[ ! -d "vault-kms-wrapper" ]]; then
  echo "Cloning vault-kms-wrapper"
  git clone git@github.com:yandex-cloud/vault-kms-wrapper.git
  cd vault-kms-wrapper
else
  echo "Vault already cloned"
  cd vault-kms-wrapper
fi

echo "Refreshing main"
git reset --hard
git checkout main
git pull
git fetch -p

if git ls-remote --exit-code origin $BRANCH; then
  echo >&2 "Remote branch '$BRANCH' already exists!"
  echo >&2 "Update suffix '$SUFFIX' to release new version"
  exit 1
fi

if git ls-remote --exit-code origin $TAG; then
  echo >&2 "Remote tag '$TAG' already exists!"
  echo >&2 "Update suffix '$SUFFIX' to release new version"
  exit 1
fi

echo "Getting github.com/hashicorp/go-kms-wrapping/v2@$VERSION"
go get github.com/hashicorp/go-kms-wrapping/v2@$VERSION
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
