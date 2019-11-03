#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
RED='\033[0;31m'
BLUE='\033[0;34m'
YELLOW='\033[0;32m'
NC='\033[0m'

function client {
  exec > >(sed "s/^/$(printf "${YELLOW}client: ${NC}")/")
  exec 2> >(sed "s/^/$(printf "${RED}client: ${NC}")/" >&2)
  ENV=debug $DIR/../bin/gourd client
}

function daemon {
  exec > >(sed "s/^/$(printf "${BLUE}daemon: ${NC}")/")
  exec 2> >(sed "s/^/$(printf "${RED}daemon: ${NC}")/" >&2)
  ENV=debug $DIR/../bin/gourdd
}

daemon & client && kill $!