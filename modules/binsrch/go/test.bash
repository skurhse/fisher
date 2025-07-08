#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o xtrace

path=$(realpath "$BASH_SOURCE")

dir=$(dirname "$path")

cd "$dir"

pushd "$dir/gen"; go generate; popd

go test -bench=.
