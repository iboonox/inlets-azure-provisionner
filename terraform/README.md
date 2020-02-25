## Inlets Azure provisionner

Terraform impletation of https://github.com/inlets/inletsctl/tree/master/pkg/provision


## Requirements

[Terraform](https://www.terraform.io/downloads.html) 0.12.x

The following Environment Variables must be set

```
- `ARM_CLIENT_ID`
- `ARM_CLIENT_SECRET`
- `ARM_SUBSCRIPTION_ID`
- `ARM_TENANT_ID`
```


## Providers

| Name | Version |
|------|---------|
| azurerm | =2.0.0 |
| random | n/a |
| template | n/a |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| location | The Azure Region in which all resources should be created. | `string` | `"westeurope"` | no |
| prefix | The prefix which should be used for all resources | `string` | `"inlets"` | no |
| vm\_size\_sku | The vm size sku name | `string` | `"Standard_B1ls"` | no |

## Outputs

| Name | Description |
|------|-------------|
| inlets\_exit\_node\_public\_ip | The public ip of the exit node |
