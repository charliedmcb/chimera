# Chimera

- [Overview](#overview)  
- [Contributing/Feedback](#contributingfeedback)
- [Dev Notes](#dev-notes)
    - [Cluster Creation](#cluster-creation)
        - [Commands](#commands)
    - [Domain Name](#domain-name)

# Overview:
This repo is the holding place of all code for [www.playchimera.net](http://www.playchimera.net) the website for understanding the chimera format, and generating decklists for online play.

# Contributing/Feedback:
This is an open source reposity maintained by charliedmcb (netrunner handle: `Santa`). I welcome anyone to open a Pull Request on the repo, and request a review which I will try to get to as timely as possible. 

For any additional feedback you can direct message me (santa181) on discord.

# Dev Notes:

## Cluster Creation:
Followed examples from this website to create the cluster, and deployment of the webserver:
https://canviz.com/hands-on-creating-a-simple-web-server-in-kubernetes/

### Commands:
1. `az group create --name ${RESOURCE_GROUP} --location ${REGION} --subscription ${AZURE_SUBSCRIPTION_ID}`
2. `az aks create --resource-group ${RESOURCE_GROUP} --name ${MC_NAME} --subscription ${AZURE_SUBSCRIPTION_ID} --node-count 1 --generate-ssh-keys --node-vm-size Standard_B2s --tier free --node-osdisk-type Ephemeral --node-osdisk-size 30 --attach-acr ${ACR_NAME}` <br>
Note: the choice of `--node-vm-size Standard_B2s`, and `--node-osdisk-type Ephemeral --node-osdisk-size 30` were made as cost cutting mesures, since the earlier version were running over budget. This might be reassessed later, depending upon the monthly cost, and resource usage/server needs/functionality <br>

## Domain Name:
Created an `App Service Domain` through the Azure Portal, partically following a mismatch of the following tutorials:
- https://learn.microsoft.com/en-us/azure/aks/app-routing?tabs=without-osm
- https://learn.microsoft.com/en-us/azure/dns/dns-getstarted-portal
- https://learn.microsoft.com/en-us/azure/app-service/manage-custom-dns-buy-domain

After the `App Service Domain` creation, a `Azure DNS Zone` was also created. I configured a routing from both `www`, and no-prefix to the external ip address exposed by the chimera golang k8s web server pod.

## Created a new nodepool:
The cluster got stopped. I believe this was caused by the subscription running out of money. I tried starting the cluster again, but it was not appearing to work. So, I created a new nodepool: <br>

`az aks nodepool add --resource-group ${RESOURCE_GROUP} --cluster-name ${MC_NAME} --name ${NODEPOOL_2} --subscription ${AZURE_SUBSCRIPTION_ID} --node-count 1 --node-vm-size Standard_B2s --node-osdisk-type Ephemeral --node-osdisk-size 30 --mode System` <br>
Note: increased to 2 nodes later.

And removed the old nodepool: <br>

`az aks nodepool delete --resource-group ${RESOURCE_GROUP} --cluster-name ${MC_NAME} --name ${NODEPOOL_1} --subscription ${AZURE_SUBSCRIPTION_ID}`

After that starting the cluster worked correctly:

`az aks start -g ${RESOURCE_GROUP} -n ${MC_NAME} --subscription ${AZURE_SUBSCRIPTION_ID}`

## Setup AKS ingress with the application routing add-on

Followed this walkthrough: <br>
https://learn.microsoft.com/en-us/azure/aks/app-routing?tabs=without-osm <br>

Including the steps to `Configure the add-on to use Azure DNS to manage DNS zones`

`openssl req -new -x509 -nodes -out aks-ingress-tls.crt -keyout aks-ingress-tls.key -subj "/CN=playchimera" -addext "subjectAltName=DNS:playchimera"`

`az keyvault create -g ${RESOURCE_GROUP} -l westus2 -n ${CHIMERA_KEYVAULT}`

`az keyvault certificate import --vault-name ${CHIMERA_KV} -n ${CHIMERA_CERT} -f aks-ingress-tls.pfx`

`az aks enable-addons -g ${RESOURCE_GROUP} -n ${MC_NAME} --addons azure-keyvault-secrets-provider,web_application_routing --enable-secret-rotation`

Note: I then later removed the AKS ingress with application routing add-on and switched back to using http, instead of https. My reasoning for this was self signing the cert caused browsers to display a massive warning when visiting the site, and having a well known source sign it is rather expensive.

