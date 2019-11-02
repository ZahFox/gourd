#!/bin/sh

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

pushd $DIR/../ >/dev/null 2>&1
go test -v ./pkg/utils
popd >/dev/null 2>&1