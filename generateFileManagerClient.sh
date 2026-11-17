#!/bin/bash
# generateFileManagerClient.sh <tag>: regenerates filemanager from
# vaud-ai-file-manager's spec at its release tag, such as v0.25.2. Open a PR
# with the result, and tag its merge filemanager/<version>.
exec "$(cd "$(dirname "$0")" && pwd)/fromTag.sh" filemanager vaud-ai-file-manager "$@"
