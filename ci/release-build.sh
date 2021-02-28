#!/usr/bin/env bash

set -e -x

mkdir -p /go/src/github.com/mt5225

ln -s  "$(pwd)/terraform-provider-gateway" "/go/src/github.com/mt5225/terraform-provider-gateway"

pushd /go/src/github.com/mt5225/terraform-provider-gateway
    ./scripts/build-release.sh
popd

pushd terraform-provider-gateway
    echo "$(git describe --abbrev=0)" > ../release-binaries/tag
    echo "Release $(git describe --abbrev=0)" > ../release-binaries/release-name
    mv pkg/*.zip ../release-binaries/
popd
