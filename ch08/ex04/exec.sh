#!/bin/sh

set -x

cd "$(dirname "$0")"

../../run reverb.go &
pid="$?"

trap "kill '${pid}'" 0               # EXIT
trap "kill '${pid}'; exit 1" 2       # INT
trap "kill '${pid}'; exit 1" 1 15    # HUP TERM

echo "hello" | ../../run ../../ch08/ex03/netcat.go localhost:8000
