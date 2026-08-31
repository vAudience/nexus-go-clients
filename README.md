# Nexus GO Clients

This repository contains the source code for the Nexus GO clients:

- **Core Client**: Interact with the [vAudience AI Core API](https://github.com/vAudience/vaud-ai-core)
- **AigentChat Client**: Interact with the [vAudience AigentChat API](https://github.com/vAudience/aigentchat)
- **FileManager Client**: Interact with the [vAudience FileManager API](https://https://github.com/vAudience/vaud-ai-file-manager)

## Installation

To install the clients, execute the following commands:

```bash
go get github.com/vaudience/nexus-go-clients/core
go get github.com/vaudience/nexus-go-clients/aigentchat
go get github.com/vaudience/nexus-go-clients/filemanager
```

## Updating Clients

To regenerate a client, run its generation script. The script will generate the client, and if there are changes, automatically commit, tag, and push:

```bash
./generateCoreClient.sh
./generateAigentChatClient.sh
./generateFileManagerClient.sh
```

The version comes from the service's own `/info/version` endpoint and is stamped into the generated package, so the package version, the git tag, and the running service always report the same version. The scripts are thin wrappers around `generateClient.sh`, which takes a package name and a service base URL.
