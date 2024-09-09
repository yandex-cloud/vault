#!/bin/bash
set -e

SCRIPT_PATH=$(dirname "${BASH_SOURCE[0]}")
. $SCRIPT_PATH/common.sh
. $SCRIPT_PATH/release_sample.cfg

init
cleanup
