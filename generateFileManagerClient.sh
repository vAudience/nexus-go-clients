#!/bin/bash
rm -rf filemanager
openapi-generator generate \
  -o filemanager \
  -i https://file-manager.dev.ai.vaud.one/info/docs/doc.json \
  -g go \
  --additional-properties packageName=filemanager  \
  --additional-properties disallowAdditionalPropertiesIfNotPresent=false \
  --git-user-id vaudience \
  --git-repo-id nexus-go-clients/filemanager
cd filemanager
rm -f .travis.yml
rm -f git_push.sh
go get github.com/stretchr/testify/assert
go mod tidy