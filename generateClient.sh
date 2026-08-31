#!/bin/bash
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$SCRIPT_DIR" || exit 1

PACKAGE="$1"
BASE_URL="$2"
if [ -z "$PACKAGE" ] || [ -z "$BASE_URL" ]; then
  echo "Usage: $0 <package-name> <base-url>"
  exit 1
fi

# The package name doubles as the output directory that gets rm -rf'd below
if ! echo "$PACKAGE" | grep -qE '^[a-z][a-z0-9_-]*$'; then
  echo "Invalid package name '${PACKAGE}'. Expected lowercase [a-z][a-z0-9_-]*."
  exit 1
fi

for BIN in curl jq go openapi-generator; do
  if ! command -v "$BIN" > /dev/null; then
    echo "Required command '${BIN}' not found. Please install it."
    exit 1
  fi
done

git pull origin main

# Fetch version from the service before generating, so it can be stamped in
VERSION=$(curl -fsS "${BASE_URL}/info/version" | jq -r '.version')
if [ -z "$VERSION" ] || [ "$VERSION" = "null" ]; then
  echo "Could not fetch version from ${BASE_URL}/info/version. Exiting."
  exit 1
fi
echo "Generating ${PACKAGE} client v${VERSION}"

rm -rf "$PACKAGE"
openapi-generator generate \
  -o "$PACKAGE" \
  -i "${BASE_URL}/info/docs/doc.json" \
  -g go \
  --additional-properties packageName="$PACKAGE" \
  --additional-properties packageVersion="$VERSION" \
  --additional-properties disallowAdditionalPropertiesIfNotPresent=false \
  --git-user-id vaudience \
  --git-repo-id "nexus-go-clients/${PACKAGE}"

cd "$PACKAGE" || exit 1
rm -f .travis.yml
rm -f git_push.sh
go get github.com/stretchr/testify/assert
go mod tidy

cd "$SCRIPT_DIR" || exit 1

# Stop if no changes
if [ -z "$(git status --porcelain)" ]; then
  echo "No changes detected. Exiting."
  exit 0
fi

# Commit, tag, push
git add -A
git commit -m "Generate ${PACKAGE} client - v${VERSION}"
git tag "${PACKAGE}/v${VERSION}"
git push origin main --tags
