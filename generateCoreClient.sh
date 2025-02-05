#!/bin/bash
rm -rf core
openapi-generator generate \
  -o core \
  -i https://core.dev.ai.vaud.one/info/docs/doc.json \
  -g go \
  --additional-properties packageName=core  \
  --additional-properties disallowAdditionalPropertiesIfNotPresent=false \
  --git-user-id vaudience \
  --git-repo-id nexus-go-clients/core
cd core
go install github.com/stretchr/testify/assert@latest
go mod tidy