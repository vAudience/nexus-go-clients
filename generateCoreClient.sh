#!/bin/bash
# generateCoreClient.sh <tag>: regenerates core from vaud-ai-core's spec at
# its release tag, such as v0.22.2. Open a PR with the result, and tag its
# merge core/<version>.
exec "$(cd "$(dirname "$0")" && pwd)/fromTag.sh" core vaud-ai-core "$@"
