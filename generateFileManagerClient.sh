#!/bin/bash
rm -rf filemanager
openapi-generator generate \
  -o filemanager \
  -i /Users/frederik/projects/vaud-ai-file-manager/docs/swagger.json \
  -g go \
  --additional-properties packageName=filemanager  \
  --additional-properties disallowAdditionalPropertiesIfNotPresent=false \
  --git-user-id vaudience \
  --git-repo-id nexus-go-clients/filemanager
cd filemanager
go get github.com/stretchr/testify/assert
go mod tidy