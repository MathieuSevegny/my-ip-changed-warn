# my-ip-changed-warn
Simple script to warn you when your IP changes

## Configuration
```docker-compose.yml
services:
  ip-warner:
    image: my-ip-changed-warn
    environment:
      - EMAIL_TO=<EMAIL_THAT_RECEIVES_WARNING>
      - EMAIL_FROM=<EMAIL_THAT_SENDS_WARNING>
      - EMAIL_TOKEN=<TOKEN_FOR_SMTP_AUTH>
      - SMTP_HOST=<SMTP_HOST>
      - DEVICE_NAME=<NAME_OF_DEVICE>
      - MAX_TRIES=10
      - SECONDS_TO_WAIT=2
      - CACHE_FOLDER_PATH=/cache
      - CACHE_FILENAME=last_ip.txt
    volumes:
      - ./cache:/cache
```