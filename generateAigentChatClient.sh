#!/bin/bash
# generateAigentChatClient.sh <tag>: regenerates aigentchat from aigentchat's
# spec at its release tag, such as v0.53.0. Open a PR with the result, and tag
# its merge aigentchat/<version>.
exec "$(cd "$(dirname "$0")" && pwd)/fromTag.sh" aigentchat aigentchat "$@"
