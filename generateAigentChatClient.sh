#!/bin/bash
rm -rf aigentchat
openapi-generator generate \
  -o aigentchat \
  -i https://aigentchat.dev.ai.vaud.one/info/docs/doc.json \
  -g go \
  --additional-properties packageName=aigentchat \
  --additional-properties disallowAdditionalPropertiesIfNotPresent=false \
  --git-user-id vaudience \
  --git-repo-id nexus-go-clients/aigentchat
cd aigentchat
go install github.com/stretchr/testify/assert@latest
go mod tidy
 