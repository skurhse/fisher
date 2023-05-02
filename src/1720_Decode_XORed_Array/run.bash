#!/usr/bin/env bash

set -o errexit

dir=$(dirname $0)
cd "$dir"

dotnet test
cd Xor.Benchmarks
dotnet run -c Release
