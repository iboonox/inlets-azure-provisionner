# Azure-Inlets

This repo is heavily inspired by [@alexellis](https://github.com/alexellis)'s cloud-provision-example as it's an implementation version of it
https://github.com/inlets/cloud-provision-example

### Architechture

Localhost > inlets proxy > exit node ( cloud ) 

RPI > inlets proxy > exit node ( cloud )

Localhost -> Edge gateway : rpi + inlets proxy > Cloud : inlets exit node ( vm )

### Authentification 

#### Environment variables based authentification

Create a service principal by running

```
az ad sp create-for-rbac -n "<yourAppName>"
```

Then set the following environment variables. See .env.tpl to see and set needed environment variables.

```
export AZURE_SUBSCRIPTION_ID=
export AZURE_TENANT_ID=
export AZURE_CLIENT_ID=
export AZURE_CLIENT_SECRET=
```

### How to run 

```
dep ensure
go run main.go --userdata-file cloud-config.txt
```

A terraform implementation [here](https://github.com/iboonox/azure-inlets/tree/master/terraform)

### TODO 
- [x] Azure provisioner initialization
- [x] Provision exit node
- [ ] Delete exit node
- [ ] List exit nodes
- [ ] Authentification enhancements ( Environment variables instead of File-based )
- [ ] Handle different scenarios : existing resource group, existing vnet ...
- [ ] Upgrade to go modules
