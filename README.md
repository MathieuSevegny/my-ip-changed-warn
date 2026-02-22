# my-ip-changed-warn
Simple service to warn you (by mail) when your public IP changes.

## Configuration (docker-compose)
```yaml
services:
  ip-warner:
    image: my-ip-changed-warn
    environment:
      - EMAIL_TO=<EMAIL_THAT_RECEIVES_WARNING> # required
      - EMAIL_FROM=<EMAIL_THAT_SENDS_WARNING> # required
      - EMAIL_TOKEN=<TOKEN_FOR_SMTP_AUTH> # required
      - SMTP_HOST=<SMTP_HOST> # required
      - DEVICE_NAME=<NAME_OF_DEVICE>
      - WAIT_TIME=2s
      - CACHE_FILE_PATH=/cache/last_ip.txt
    volumes:
      - ./cache:/cache
```