#!/bin/bash
exec "$(cd "$(dirname "$0")" && pwd)/generateClient.sh" \
  core https://core.dev.ai.vaud.one
