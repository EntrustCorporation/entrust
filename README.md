
# Golang client for the Entrust Certificate Services Enterprise API

This is a simple and (currently) only partial implementation of the API.

**This project is made available by Entrust Corporation under the [MIT license](LICENSE).** As such the client is provided to you AS IS and without warranty, as further detailed in the MIT license.

## Config

| Environment variable            | Value                            |
|---------------------------------|----------------------------------|
| `ENTRUST_API_CERTIFICATE`       | PEM encoded client certificate   |
| `ENTRUST_API_CERTIFICATE_PATH`  | Client certificate path          |
| `ENTRUST_API_PRIVATE_KEY`       | PEM encoded private key          |
| `ENTRUST_API_PRIVATE_KEY_PATH`  | Private key path                 |
| `ENTRUST_API_USERNAME`          | API Username                     |
| `ENTRUST_API_PASSWORD`          | API Password (API Key)           |

The environment variables `ENTRUST_API_CERTIFICATE` and `ENTRUST_API_PRIVATE_KEY` take presedence over `ENTRUST_API_CERTIFICATE_PATH` and `ENTRUST_API_PRIVATE_KEY_PATH`.

## Obtaining credentials

These credentials can be obtained via the Entrust Certificate Services Enterprise Portal (Enterprise Portal, Administration > Advanced Settings > API). In order to obtain an account and related credentials, you must have an active subscription or other entitlement from Entrust for Entrust Certificate Services (ECS).

# Security

The Entrust credentials provide strong entitlements and can have a devastating impact to your organisation or service when not secured properly. It’s your sole responsibility to store these credentials securely and restrict API credentials to the resources and capabilities as required.

**Do not put any credential directly in your configuration file!**