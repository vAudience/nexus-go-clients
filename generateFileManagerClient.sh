#!/bin/bash
exec "$(cd "$(dirname "$0")" && pwd)/generateClient.sh" \
  filemanager https://file-manager.dev.ai.vaud.one
