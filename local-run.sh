#!/bin/bash

# TODO: move build to a separate build.sh

set -e

# TODO: go generate

mkdir -p bin
go build -o ./bin/site ./cmd/site
bin/site
