# 18. Azure CLI

## Commands used during the demo
```bash
az vm list-skus --size Standard_DSv5 --all --output table --location polandcentral
```

```bash
az vm create --resource-group ca-devopsua6-01 --name ca-test-01 --image Ubuntu2404 --ssh-key-values ~/.ssh/id_ed25519.pub --size Standard_D2_v5
```
```bash
az group --help
```

```bash
az group list
```
```bash
az network nsg list
```

```bash
az network
```

```bash
az network --help
```
```bash
az network public-ip list
```
```bash
az vm delete --name ca-test-01 --resource-group ca-devopsua6-01
```

## Additional tasks 

Create a script to provision n-number of machines (not more then 3). The script should: 
* Have a usage menu
* Have restrictions on the size of the machine (3 types max)
* Ability tu enter the number of VMs to provition.
* VM should be provisioned in the same Virtual Network. 
* Have the ability to delete provisioned VMs.