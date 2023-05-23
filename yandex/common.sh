init() {
  START_DIR=$(pwd)
  trap 'cd $START_DIR' EXIT

  SCRIPT_DIR=$(dirname "${BASH_SOURCE[0]}")
  cd $SCRIPT_DIR
  SCRIPT_DIR=$(pwd)

  if [[ -n $WORK_DIR ]]; then
    mkdir -p $WORK_DIR
    cd $WORK_DIR
  fi
  WORK_DIR=$(pwd)
  YCKMS_VERSION="$BASE_VERSION+yckms"
}

go_to_work_dir() {
  cd $WORK_DIR
}

cleanup() {
  go_to_work_dir
  rm -rf vault
  rm -rf vault-kms-wrapper
}

init_vault() {
  go_to_work_dir
  if [[ ! -d "vault" ]]; then
    echo "Cloning vault"
    git clone git@github.com:yandex-cloud/vault.git
    cd vault
    git remote add upstream git@github.com:hashicorp/vault.git
  else
    echo "Vault already cloned"
    cd vault
    git reset --hard
  fi

  echo "Synchronizing vault with upstream"
  git checkout main
  git pull upstream main
  echo "Fetching tags"
  git fetch upstream --tags

  go_to_work_dir
}

init_vault_kms_wrapper() {
  go_to_work_dir
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

  go_to_work_dir
}

get_kms_wrapper_version() {
  go_to_work_dir
  cd vault
  CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
  git checkout $BASE_VERSION
  KMS_WRAPPER_VERSION=$(go list -m github.com/hashicorp/go-kms-wrapping/v2 | cut -f 2 -d " ")
  echo "Current go-kms-wrapping version: $KMS_WRAPPER_VERSION"
  git checkout $CURRENT_BRANCH
  go_to_work_dir
}
