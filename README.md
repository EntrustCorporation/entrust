
# Golang client for the Entrust Certificate Services Enterprise API

This is a simple and (currently) only partial implementation of the API.

## Config

| Environment variable            | Value                            |
|---------------------------------|----------------------------------|
| `ENTRUST_API_CERTIFICATE`       | PEM encoded client certificate   |
| `ENTRUST_API_CERTIFICATE_FILE`  | Client certificate filename      |
| `ENTRUST_API_PRIVATE_KEY`       | PEM encoded private key          |
| `ENTRUST_API_PRIVATE_KEY_FILE`  | Private key filename             |
| `ENTRUST_API_USERNAME`          | API Username                     |
| `ENTRUST_API_PASSEWORD`         | API Password (API Key)           |

The environment variables `ENTRUST_API_CERTIFICATE` and `ENTRUST_API_PRIVATE_KEY` take presedence over `ENTRUST_API_CERTIFICATE_FILE` and `ENTRUST_API_PRIVATE_KEY_FILE`.

## Obtaining credentials

These credentials can be obtained via the Enterprise Portal, `Administration > Advanced Settings > API` and your account must be enabled for API usage.