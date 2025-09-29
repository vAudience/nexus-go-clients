# Nexus GO Clients

This repository contains the source code for the Nexus GO clients. Currently, two clients are available:

- **Core Client**: Interact with the [vAudience AI Core API](https://github.com/vAudience/vaud-ai-core)
- **AigentChat Client**: Interact with the [AigentChat API](https://github.com/vAudience/aigentchat)
- **FileManager Client**: Interact with the [vAudience File Manager API](https://https://github.com/vAudience/vaud-ai-file-manager)

## Installation

To install the clients, execute the following commands:

```bash
go get github.com/vaudience/nexus-go-clients/core
go get github.com/vaudience/nexus-go-clients/aigentchat
go get github.com/vaudience/nexus-go-clients/filemanager
```

## Updating Clients

To update the clients, execute the following commands:

```bash
./generateCoreClient.sh
./generateAigentChatClient.sh
./generateFileManagerClient.sh
```

## Tags

To create tags, use the following commands:

```bash
git tag core/v0.1.0
git tag aigentchat/v0.1.0
git tag filemanager/v0.1.0
```
