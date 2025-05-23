![release](https://github.com/DelineaXPM/terraform-provider-dsv/workflows/release/badge.svg)

# Delinea DevOps Secrets Vault - Terraform Provider

The [Delinea](https://delinea.com/) [DevOps Secrets Vault](https://delinea.com/products/devops-secrets-management-vault) (DSV) [Terraform](https://www.terraform.io/) Provider makes Secrets data available and provisions client secrets for existing roles.

## Installation

The latest release can be downloaded from [here](https://github.com/DelineaXPM/terraform-provider-dsv/releases/latest).

### Terraform 0.12 and earlier

Extract the specific file for your OS and Architecture to the plugins directory of the user's profile. You may have to create the directory.

| OS      | Default Path                    |
| ------- | ------------------------------- |
| Linux   | `~/.terraform.d/plugins`        |
| Windows | `%APPDATA%\terraform.d\plugins` |

### Terraform 0.13 and later

Terraform 0.13 uses a different file system layout for 3rd party providers. More information on this can be found [here](https://www.terraform.io/upgrade-guides/0-13.html#new-filesystem-layout-for-local-copies-of-providers). The following folder path will need to be created in the plugins directory of the user's profile.

#### Windows

```text
%APPDATA%\TERRAFORM.D\PLUGINS
└───terraform.delinea.com
    └───delinea
        └───dsv
            └───1.0.0
                └───windows_amd64
```

#### Linux

```text
~/.terraform.d/plugins
└───terraform.delinea.com
    └───delinea
        └───dsv
            └───1.0.0
                ├───linux_amd64
```

## Usage

For Terraform 0.13+, include the `terraform` block in your configuration or plan to that specifies the provider:

```terraform
terraform {
    required_providers {
        dsv = {
            source = "terraform.delinea.com/delinea/dsv"
            version = "~> 1.0"
        }
    }
}
```

To run the example, create a `terraform.tfvars`:

```hcl
dsv_client_id     = ""
dsv_client_secret = ""
dsv_tenant        = "mytenant"
dsv_role_name     = "example-role"
dsv_secret_path   = "/path/to/a/test/secret"
```

To run with AWS as auth provider

```hcl
dsv_auth_provider = "aws"
dsv_tenant        = "mytenant"
dsv_role_name     = "example-role"
dsv_secret_path   = "/path/to/a/test/secret"
```

To run with Azure as auth provider

```hcl
dsv_auth_provider = "azure"
dsv_tenant        = "mytenant"
dsv_role_name     = "example-role"
dsv_secret_path   = "/path/to/a/test/secret"
```

## Troubleshooting

### Failure running `terraform init -upgrade`

- Try pinning the provider version if you are currently using a range like `~> 1.0` and instead use the latest version found on the [registry](https://registry.terraform.io/providers/DelineaXPM/dsv/latest).
- Additionally, if you have `TF_PLUGIN_CACHE_DIR` environment set, try `unset TF_PLUGIN_CACHE_DIR` or `$ENV:TF_PLUGIN_CACHE_DIR = $null` and run the command again to see if it helps the upgrade succeed.
