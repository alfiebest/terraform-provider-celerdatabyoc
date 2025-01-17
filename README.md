# CelerData Cloud BYOC provider

You can use the CelerData Cloud BYOC provider to interact with most CelerData Cloud BYOC resources. By default, you have the required AWS skills before you can follow the tutorials in this article.

### What you should prepare

1. A [CelerData Cloud BYOC](https://cloud.celerdata.com/login) account
2. You need to navigate to `Application keys` in CelerData Cloud BYOC console, click `+ New a secret` in `Application keys`, and then find `Client_id` , `Secret` information.

### Initialize provider

All of the CelerData Cloud BYOC terraform provider resources can be created in a dedicated terraform module for your environment.

```
terraform {
  required_providers {
    celerdatabyoc = {
      source = "CelerData/celerdatabyoc"
      version = "[provider version]"
    }
  }
}

provider "celerdatabyoc" {
  client_id = "[your client id]"
  client_secret = "[your client secret]"
}

```
