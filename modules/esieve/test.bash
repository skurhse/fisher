#!/usr/bin/env bash

# REQ: Runs go tests and benchmarks. <eris 2023-05-11>

set -Cefuvxo pipefail

dir=$(dirname "$0")
cd "$dir"

go test -v -bench=. -benchmem
