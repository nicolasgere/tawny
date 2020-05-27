# Install and configure Tawny


## Installation

For now, only a docker image is supported as deliverable.

```
docker pull elmazout/tawny:latest
```



## Configuration

Tawny can be configured using environment variable

| Variable | Description | Default
| --- | --- | --- |
| TAWNY_BADGER_PATH | Path for badger to store data | db-temp/badger 
| TAWNY_HTTPS_ENABLE | Enable https | false
| TAWNY_HTTPS_DOMAIN | if TAWNY_HTTPS_ENABLE is true then TAWNY_HTTPS_DOMAIN is required  | 
| TAWNY_HTTPS_EMAIL | if TAWNY_HTTPS_ENABLE is true then TAWNY_HTTPS_EMAIL is required |
| TAWNY_HTTPS_PORT | port used by web gateway to listen for https | 443
| TAWNY_HTTP_PORT | port used by web gateway to listen for http | 80
| TAWNY_GRPC_PORT| Show file differences that haven't been staged | 4000