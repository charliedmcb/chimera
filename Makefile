VERSION ?= v0.0.48-alpha
DEPLOYMENT_VERSION ?= $(subst .,-,$(VERSION))

web:
	cd deck-builder && go run .

codegen: 
	cd hack/codegen && go run .

login:
	az login

build:
	go build -o ./bin/webapp ./deck-builder/
	docker build --tag $(ACR_NAME).azurecr.io/$(ACR_REPOSITORY_NAME):$(VERSION) . --no-cache
	rm -rf ./bin

publish:
	az acr login --name $(ACR_NAME) --resource-group $(RESOURCE_GROUP) --subscription $(SUBSCRIPTION_ID)
	TOKEN=$(az acr login --name $(ACR_NAME) --subscription $(SUBSCRIPTION_ID) --expose-token --output tsv --query accessToken) \
	$(shell docker login ${ACR_NAME}.azurecr.io --username 00000000-0000-0000-0000-000000000000 --password-stdin <<< ${TOKEN})

	docker push $(ACR_NAME).azurecr.io/$(ACR_REPOSITORY_NAME):$(VERSION)

release:
	az aks get-credentials --resource-group $(RESOURCE_GROUP) --subscription $(SUBSCRIPTION_ID) --name $(MC_NAME)
	kubectl create deployment deck-builder-$(DEPLOYMENT_VERSION) --image=$(ACR_NAME).azurecr.io/$(ACR_REPOSITORY_NAME):$(VERSION) --port=80
	kubectl expose deployment deck-builder-$(DEPLOYMENT_VERSION) --port=80 --type=LoadBalancer

cleanup:
	az aks get-credentials --resource-group $(RESOURCE_GROUP) --subscription $(SUBSCRIPTION_ID) --name $(MC_NAME)
	kubectl delete deployment deck-builder-$(DEPLOYMENT_VERSION)
	kubectl delete service deck-builder-$(DEPLOYMENT_VERSION)

deploy: login build publish release