#!/usr/bin/env bash

set -e -x

cat terraform-provider-gateway/.git/ref

mkdir -p /go/src/github.com/mt5225

ln -s  "$(pwd)/terraform-provider-gateway" "/go/src/github.com/mt5225/terraform-provider-gateway"

pushd /go/src/github.com/mt5225/terraform-provider-gateway
    make test
    make testacc
    make vet
popd
