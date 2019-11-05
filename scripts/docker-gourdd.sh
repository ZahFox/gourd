#!/bin/sh

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
HASIMG=$(docker images "gourd" | wc -l)

if [[ "$HASIMG" -lt "2" ]]
then
  pushd $DIR/../
  docker build --force-rm -t gourd .
  popd
fi

CID=$(docker run -d --rm --name=gourd gourd)

docker exec -it gourd /bin/sh

echo "you can remove the gourd container with the following command:"
echo "docker rm -f $CID"