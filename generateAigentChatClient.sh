#!/bin/bash
exec "$(cd "$(dirname "$0")" && pwd)/generateClient.sh" \
  aigentchat https://aigentchat.dev.ai.vaud.one
