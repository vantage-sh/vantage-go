#!/usr/bin/env bash
set -e
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
pushd $SCRIPT_DIR
curl https://api.vantage.sh/v2/swagger.json | jq . > swagger.json
swagger generate client -f swagger.json -c vantage vantagev2
go mod tidy
popd

