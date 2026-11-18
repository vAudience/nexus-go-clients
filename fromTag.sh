#!/bin/bash
# fromTag.sh <part> <service repo> <tag>
#
# Regenerates <part> from the service's committed spec, docs/swagger.json, at
# its release tag, through gh (logged in to GitHub), with the tag's version.
set -euo pipefail
if [ $# -ne 3 ] || [ -z "$3" ]; then
  echo "Usage: $0 <part> <service repo> <tag>, e.g. $0 core vaud-ai-core v0.22.1" >&2
  exit 1
fi
PART="$1"
SERVICE="$2"
TAG="$3"
if ! [[ "$TAG" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  echo "Invalid tag '${TAG}': expected the service's release tag, vX.Y.Z." >&2
  exit 1
fi
SPEC="$(mktemp)"
trap 'rm -f "$SPEC"' EXIT
if ! gh api -H 'Accept: application/vnd.github.raw' "repos/vAudience/${SERVICE}/contents/docs/swagger.json?ref=${TAG}" > "$SPEC"; then
  echo "Could not read ${SERVICE}'s docs/swagger.json at ${TAG}." >&2
  exit 1
fi
"$(cd "$(dirname "$0")" && pwd)/generateClient.sh" "$PART" "$SPEC" "$TAG"
