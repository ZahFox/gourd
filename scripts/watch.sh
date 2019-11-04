#!/bin/bash

CMD=install

if [ ! -z "$1" ]
then
    CMD=$1
fi

reflex -r '\.go$' make $CMD