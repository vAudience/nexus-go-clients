#!/bin/bash
# check.sh: vets, builds and tests every generated part, each a module of its own.
set -euo pipefail
cd "$(dirname "$0")"
for part in */; do
  part="${part%/}"
  [ -f "$part/go.mod" ] || continue
  echo "== ${part}"
  (cd "$part" && go vet ./... && go build ./... && go test ./...)
done
