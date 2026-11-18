#!/bin/bash
# generateClient.sh <part> <spec file> <version>
#
# Generates the Go client <part>/ from a service's API spec, stamped with the
# service's version, and tidies its module. It replaces <part>/ entirely and
# never commits, tags or pushes: the change goes through a PR, and a person
# tags its merge <part>/v<version>.
set -euo pipefail

# Every part was generated with this openapi-generator: another version would
# rewrite whole clients for nothing. OPENAPI_GENERATOR names another command
# that runs it, e.g. "npx --yes @openapitools/openapi-generator-cli", which reads the
# version from openapitools.json.
GENERATOR_VERSION=7.23.0
read -r -a GENERATOR <<< "${OPENAPI_GENERATOR:-openapi-generator}"

usage() {
  echo "Usage: $0 <part> <spec file> <version>, e.g. $0 core ./swagger.json 0.22.1" >&2
  exit 1
}
[ $# -eq 3 ] || usage
PART="$1"
VERSION="${3#v}"

# The part doubles as the directory removed below and as the Go package
# clause, where a hyphen would not compile.
if ! [[ "$PART" =~ ^[a-z][a-z0-9_]*$ ]]; then
  echo "Invalid part '${PART}': expected lowercase [a-z][a-z0-9_]*." >&2
  exit 1
fi
if ! [[ "$VERSION" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  echo "Invalid version '$3': expected X.Y.Z, the service's release tag without its v." >&2
  exit 1
fi
if [ ! -f "$2" ]; then
  echo "No spec file at '$2'." >&2
  exit 1
fi
SPEC="$(cd "$(dirname "$2")" && pwd -P)/$(basename "$2")"

# The rest runs in the repository, so the version check and the generation
# read the same openapitools.json, also through npx.
ROOT="$(cd "$(dirname "$0")" && pwd -P)"
cd "$ROOT"

# <part>/ is replaced below, so the spec can't be inside it.
case "$SPEC" in
  "$ROOT/$PART"/*)
    echo "The spec '$2' is inside ${PART}/, which is replaced: copy it out of the part first." >&2
    exit 1
    ;;
esac

for BIN in go "${GENERATOR[0]}"; do
  if ! command -v "$BIN" > /dev/null; then
    echo "Required command '${BIN}' not found." >&2
    exit 1
  fi
done
FOUND="$("${GENERATOR[@]}" version 2> /dev/null | tail -n 1)"
if [ "$FOUND" != "$GENERATOR_VERSION" ]; then
  echo "openapi-generator ${GENERATOR_VERSION} is needed, '${GENERATOR[*]}' is ${FOUND:-not runnable}: see the README." >&2
  exit 1
fi

echo "Generating ${PART} v${VERSION} from ${SPEC}"
rm -rf "$PART"
"${GENERATOR[@]}" generate \
  -o "$PART" \
  -i "$SPEC" \
  -g go \
  --additional-properties packageName="$PART" \
  --additional-properties packageVersion="$VERSION" \
  --additional-properties disallowAdditionalPropertiesIfNotPresent=false \
  --git-user-id vaudience \
  --git-repo-id "nexus-go-clients/${PART}"

cd "$PART"
rm -f .travis.yml git_push.sh
go get github.com/stretchr/testify/assert
go mod tidy
echo "Generated ${PART} v${VERSION}: check it with ./check.sh, then open a PR; tag its merge ${PART}/v${VERSION}."
