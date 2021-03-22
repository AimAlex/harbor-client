#!/bin/bash

set -e

if !(command -v swagger >/dev/null 2>&1); then
  dir=$(mktemp -d)
  git clone https://github.com/go-swagger/go-swagger "$dir"
  pushd "$dir" > /dev/null || return
  git checkout v0.25.0
  go install -ldflags "-X github.com/go-swagger/go-swagger/cmd/swagger/commands.Version=$(git describe --tags) -X github.com/go-swagger/go-swagger/cmd/swagger/commands.Commit=$(git rev-parse HEAD)" ./cmd/swagger
  popd || return
fi
# mkdir -p ./harborcli
swagger generate client -f ./doc/swagger.yaml -c harborcli -t ./