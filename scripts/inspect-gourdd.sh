#!/bin/sh

trap ctrl_c INT
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

SDIR=/tmp/.gourd/
SNAM=gourdd-debug.sock
SPATH="${SDIR}${SNAM}"

DPID=0

function ctrl_c {
  kill $DPID
}

function daemon {
  ENV=debug $DIR/../bin/gourdd &
  DPID=$!
  mv $SPATH $SPATH.original
}

function monitor {
	echo "Monitoring gourdd at: $SPATH"
	socat -t100 -x -v UNIX-LISTEN:$SPATH,mode=777,reuseaddr,fork UNIX-CONNECT:$SPATH.original
}

daemon && monitor && kill $DPID 
