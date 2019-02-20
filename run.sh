#!/bin/bash
trap "kill 0" EXIT

if [ -z "$GOPATH" ]; then
      echo "\$GOPATH variable must be set."
      exit 1
fi

if [ ! -f $GOPATH/bin/server ]; then
    echo "Could not find server executable file, run 'test_and_build.sh'"
    exit 1
fi

if [ ! -f $GOPATH/bin/client ]; then
    echo "Could not find client executable file, run 'test_and_build.sh'"
    exit 1
fi

$GOPATH/bin/server &
$GOPATH/bin/client &

wait