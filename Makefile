PROJECT_NAME := Pulumi K3OS Resource Provider

PACK             := k3os
PACKDIR          := sdk
PROJECT          := github.com/spigell/pulumi-k3os
NODE_MODULE_NAME := @spigell/pulumi-k3os

PROVIDER        := pulumi-resource-${PACK}
CODEGEN         := pulumi-gen-${PACK}
PROVIDER_PATH   := provider
VERSION_PATH     := ${PROVIDER_PATH}/pkg/version.Version

SCHEMA_FILE     := provider/cmd/pulumi-resource-k3os/schema.json
GOPATH			:= $(shell go env GOPATH)

WORKING_DIR     := $(shell pwd)
TESTPARALLELISM := 4

ensure::
	cd provider && go mod tidy
	cd sdk && go mod tidy

gen::
	(cd provider && go build -o $(WORKING_DIR)/bin/${CODEGEN} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/$(CODEGEN))

provider_codegen:
	(cd provider && VERSION=${VERSION} go generate cmd/${PROVIDER}/main.go)

provider:: provider_codegen
	(cd provider && go build -o $(WORKING_DIR)/bin/${PROVIDER} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" $(PROJECT)/${PROVIDER_PATH}/cmd/$(PROVIDER))

provider_debug::
	(cd provider && go build -o $(WORKING_DIR)/bin/${PROVIDER} -gcflags="all=-N -l" -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" $(PROJECT)/${PROVIDER_PATH}/cmd/$(PROVIDER))

go_sdk::
	rm -rf sdk/go
	$(WORKING_DIR)/bin/$(CODEGEN) -version=${VERSION} go $(SCHEMA_FILE) $(CURDIR)

nodejs_sdk::
	rm -rf sdk/nodejs
	$(WORKING_DIR)/bin/$(CODEGEN) -version=${VERSION} nodejs $(SCHEMA_FILE) $(CURDIR)
	cd ${PACKDIR}/nodejs/ && \
		yarn install && \
		yarn run build
	cp README.md LICENSE ${PACKDIR}/nodejs/package.json ${PACKDIR}/nodejs/yarn.lock ${PACKDIR}/nodejs/bin/
	sed -i.bak 's/$${VERSION}/$(VERSION)/g' ${PACKDIR}/nodejs/bin/package.json

python_sdk::
	$(WORKING_DIR)/bin/$(CODEGEN) -version=${VERSION} python $(SCHEMA_FILE) $(CURDIR)
	cp README.md ${PACKDIR}/python/
	cd ${PACKDIR}/python/ && \
		python3 setup.py clean --all 2>/dev/null && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		sed -i.bak -e 's/^VERSION = .*/VERSION = "$(PYPI_VERSION)"/g' -e 's/^PLUGIN_VERSION = .*/PLUGIN_VERSION = "$(VERSION)"/g' ./bin/setup.py && \
		rm ./bin/setup.py.bak && \
		cd ./bin && python3 setup.py build sdist

.PHONY: build
build:: gen provider go_sdk nodejs_sdk python_sdk

# Required for the codegen action that runs in pulumi/pulumi
only_build:: build

lint::
	for DIR in "provider" "sdk"; do \
		cd $$DIR && golangci-lint run -c ../.golangci.yml --timeout 10m ; cd -; \
	done


install:: install_nodejs_sdk
	cp $(WORKING_DIR)/bin/${PROVIDER} ${GOPATH}/bin


install_python_sdk::
	#target intentionally blank

install_go_sdk::
	#target intentionally blank

install_nodejs_sdk::
	-yarn unlink --cwd $(WORKING_DIR)/sdk/nodejs/bin
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin
