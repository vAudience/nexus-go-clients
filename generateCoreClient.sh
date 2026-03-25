#!/bin/bash
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

git pull origin main

rm -rf core
openapi-generator generate \
  -o core \
  -i https://core.dev.ai.vaud.one/info/docs/doc.json \
  -g go \
  --additional-properties packageName=core  \
  --additional-properties disallowAdditionalPropertiesIfNotPresent=false \
  --git-user-id vaudience \
  --git-repo-id nexus-go-clients/core
cd core
rm -f .travis.yml
rm -f git_push.sh
go get github.com/stretchr/testify/assert
go mod tidy

cd "$SCRIPT_DIR"

# Stop if no changes
if [ -z "$(git status --porcelain)" ]; then
  echo "No changes detected. Exiting."
  exit 0
fi

# Extract version from generated README.md
VERSION=$(sed -n 's/^- API version: \([0-9][0-9]*\.[0-9][0-9]*\.[0-9][0-9]*\)/\1/p' core/README.md)
if [ -z "$VERSION" ]; then
  echo "Could not extract version from README.md. Exiting."
  exit 1
fi

# Commit, tag, push
git add -A
git commit -m "Generate core client - v${VERSION}"
git tag "core/v${VERSION}"
git push origin main --tags
