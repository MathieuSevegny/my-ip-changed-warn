# my-ip-changed-warn
Simple service to warn you (by mail) when your public IP changes. 

Uses https://api.ipify.org/ by default to fetch public IP (configurable).

## Environment variables

| Name           | Description                                                                                     | Required | Default value             |
|----------------|-------------------------------------------------------------------------------------------------|----------|---------------------------|
| EMAIL_TO       | Email address that will receive the warning email.                                                 | Yes      |                           |
| EMAIL_FROM     | Email address that will be used as sender of the warning email.                                                 | Yes      |                           |
| EMAIL_TOKEN    | Token that will be used for SMTP authentication.                                                 | Yes      |
| SMTP_HOST      | Host of the SMTP server that will be used to send the warning email.                                                 | Yes      |                           |
| API_ENDPOINT   | Endpoint to fetch the public IP.                                                 | No       | https://api.ipify.org/   |
| DEVICE_NAME    | Name of the device that will be included in the warning email.                                                 | No       | your server                          |
| WAIT_TIME      | Time to wait between each check for IP change.                                                 | No       | 5m                        |
| DATA_FILE_PATH | Path to the file where the last known IP will be stored.                                                 | No       | /data/last_ip.txt         |

## Docker-Compose configuration
```yaml
services:
  ip-warner:
    image: mathieusevegny/my-ip-changed-warn
    environment:
      - EMAIL_TO=<EMAIL_THAT_RECEIVES_WARNING> # required
      - EMAIL_FROM=<EMAIL_THAT_SENDS_WARNING> # required
      - EMAIL_TOKEN=<TOKEN_FOR_SMTP_AUTH> # required
      - SMTP_HOST=<SMTP_HOST> # required
      - API_ENDPOINT=https://api.ipify.org/
      - DEVICE_NAME=<NAME_OF_DEVICE>
      - WAIT_TIME=5m
      - DATA_FILE_PATH=/data/last_ip.txt
    volumes:
      - ./tmp:/data
```