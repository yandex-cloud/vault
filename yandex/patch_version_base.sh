#!/bin/bash
set -e

START_DIR=$(pwd)
trap 'cd $START_DIR' EXIT

SCRIPT_PATH=$(dirname "${BASH_SOURCE[0]}")
cd "$SCRIPT_PATH"/..

sed -i '' 's/.*VersionMetadata.*=.*""/VersionMetadata = "yckms"/' version/version_base.go
go fmt version/version_base.go
