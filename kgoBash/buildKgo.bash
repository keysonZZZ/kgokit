#!/usr/bin/env bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $DIR/..
export GOPATH=`pwd`
go install github.com/keysonZZZ/kgo/kgo

