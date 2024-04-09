#!/usr/bin/env bash
set -e
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
pushd $SCRIPT_DIR
curl "${VANTAGE_HOST:-https://api.vantage.sh}/v1/swagger.json" | jq . > swagger.json
swagger generate client -f swagger.json -c vantage vantagev1
popd

