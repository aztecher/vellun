PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
IMG ?= localhost/controller:latest

.PHONY: all
all: help

##@ General

.PHONY: help
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

# TODO: manifest
# TODO: generate

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: lint
lint: golangci-lint ## Run golangci-lint linter
	$(GOLANGCI_LINT) run

.PHONY: lint-fix
lint-fix: golangci-lint ## Run golangci-lint linter and perform fixes
	$(GOLANGCI_LINT) run --fix

.PHONY: lint-config
lint-config: golangci-lint ## Verify golangci-lint linter configuration
	$(GOLANGCI_LINT) config verify

.PHONY: test
test: fmt vet setup-envtest ## Run tests.
	KUBEBUILDER_ASSETS="$(shell $(ENVTEST) use $(ENVTEST_K8S_VERSION) --bin-dir $(LOCALBIN) -p path)" go test $$(go list ./... | grep -v /e2e) -coverprofile cover.out

.PHONY: test
test-verbose: manifests generate fmt vet setup-envtest ## Run tests.
	KUBEBUILDER_ASSETS="$(shell $(ENVTEST) use $(ENVTEST_K8S_VERSION) --bin-dir $(LOCALBIN) -p path)" go test $$(go list ./... | grep -v /e2e) -coverprofile cover.out -v -- --nocapture

.PHONY: test-e2e
test-e2e: fmt vet ## Run the e2e tests. Expected an isolated environment using Kind.
	@command -v kind >/dev/null 2>&1 || { \
		echo "Kind is not installed. Please install Kind manually."; \
		exit 1; \
	}
	@kind get clusters | grep -q 'kind' || { \
		echo "No Kind cluster is running. Please start a Kind cluster before running the e2e tests."; \
		exit 1; \
	}
	go test ./test/e2e/ -v -ginkgo.v

##@ Build

.PHONY: build
build: manifests generate fmt vet ## Build manager binary.
	go build -o bin/manager cmd/main.go

.PHONY: run
run: manifests generate fmt vet ## Run a controller from your host.
	go run ./cmd/main.go

# If you wish to build the manager image targeting other platforms you can use the --platform flag.
# (i.e. docker build --platform linux/arm64). However, you must enable docker buildKit for it.
# More info: https://docs.docker.com/develop/develop-images/build_enhancements/
.PHONY: docker-build
docker-build: ## Build docker image with the manager.
	$(CONTAINER_TOOL) build -t ${IMG} .

#TODO: docker-buildx
#TODO: build-installer

##@ Deployment

ifndef ignore-not-found
  ignore-not-found = false 
endif

.PHONY: install
install: kustomize ## Install CRDs into the K8s cluster specified in ~/.kube/config.
	$(KUSTOMIZE) build config/crd | $(KUBECTL) apply -f -

.PHONY: uninstall
uninstall: kustomize ## Uninstall CRDs from the K8s cluster specified in ~/.kube/config. Call with ignor
	$(KUSTOMIZE) build config/crd | $(KUBECTL) delete --ignore-not-found=$(ignore-not-found) -f -

.PHONY: deploy
deploy: kustomize ## Deploy controller to the K8s cluster specified in ~/.kube/config.
	cd config/manager && $(KUSTOMIZE) edit set image controller=${IMG}
	$(KUSTOMIZE) build config/default | $(KUBECTL) apply -f -

undeploy: kustomize ## Undeploy controller from the K8s cluster specified in ~/.kube/config. Call with ignore-not-found=true to ignore resource not found errors during deletion.
	$(KUSTOMIZE) build config/default | $(KUBECTL) delete --ignore-not-found=$(ignore-not-found) -f -

##@ Tools

## Location to install dependencies to

LOCALBIN ?= $(PROJECT_DIR)/bin
$(LOCALBIN):
	@mkdir -p $(LOCALBIN)

## Tool Binaries
KIND ?= $(LOCALBIN)/kind
KUBECTL ?= $(LOCALBIN)/kubectl
KUBEBUILDER ?= $(LOCALBIN)/kubebuilder
DOCKERBINDIR ?= /usr/bin/
DOCKER ?= $(DOCKERBINDIR)/docker
KUSTOMIZE ?= $(LOCALBIN)/kustomize
CONTROLLER_GEN ?= $(LOCALBIN)/controller-gen
ENVTEST ?= $(LOCALBIN)/setup-envtest
GOLANGCI_LINT ?= $(LOCALBIN)/golangci-lint

CONTAINER_TOOL ?= $(DOCKER)

## Tool Variables
KIND_PREFIX ?= velocis
KIND_CLUSTER_NAME ?= $(KIND_PREFIX)

### Tool Versions

KIND_VERSION ?= v0.27.0
KIND_BIN_URL ?= https://kind.sigs.k8s.io/dl/$(KIND_VERSION)/kind-linux-amd64

KUBECTL_VERSION ?= v1.32.2
KUBECTL_BIN_URL ?= https://dl.k8s.io/release/$(KUBECTL_VERSION)/bin/linux/amd64/kubectl

KUBEBUILDER_VERSION ?= latest
KUBEBUILDER_BIN_URL ?= https://go.kubebuilder.io/dl/$(KUBEBUILDER_VERSION)/linux/amd64

KUSTOMIZE_VERSION ?= v5.5.0
KUSTOMIZE_BIN_URL ?= sigs.k8s.io/kustomize/kustomize/v5

CONTROLLER_TOOLS_VERSION ?= v0.17.2
CONTROLLER_GEN_BIN_URL ?= sigs.k8s.io/controller-tools/cmd/controller-gen

ENVTEST_VERSION ?= release-0.20
ENVTEST_BIN_URL ?= sigs.k8s.io/controller-runtime/tools/setup-envtest

ENVTEST_K8S_VERSION ?= 1.32
ENVTEST_K8S_VERSION ?= release-0.20

GOLANGCI_LINT_VERSION ?= v1.63.4
GOLANGCI_LINT_BIN_URL ?= github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: kind
kind: $(KIND) ## Download kind locally if necessary
$(KIND): | $(LOCALBIN)
	@curl -Lo $(KIND) $(KIND_BIN_URL)
	@chmod +x $(KIND)

.PHONY: kubectl
kubectl: $(KUBECTL) ## Download kubectl locally if necessary
$(KUBECTL): | $(LOCALBIN)
	@curl -Lo $(KUBECTL) $(KUBECTL_BIN_URL)
	@chmod +x $(KUBECTL)

.PHONY: kubebuilder
kubebuilder: $(KUBEBUILDER) ## Download kubebuilder locally if necessary
$(KUBEBUILDER): | $(LOCALBIN)
	@curl -Lo $(KUBEBUILDER) $(KUBEBUILDER_BIN_URL)
	@chmod +x $(KUBEBUILDER)

.PHONY: kustomize
kustomize: $(KUSTOMIZE) ## Download kustomize locally if necessary
$(KUSTOMIZE): $(LOCALBIN)
	$(call go-install-tool,$(KUSTOMIZE),$(KUSTOMIZE_BIN_URL),$(KUSTOMIZE_VERSION))

.PHONY: controller-gen
controller-gen: $(CONTROLLER_GEN) ## Download controller-gen locally if necessary.
$(CONTROLLER_GEN): $(LOCALBIN)
	$(call go-install-tool,$(CONTROLLER_GEN),$(CONTROLLER_GEN_BIN_URL),$(CONTROLLER_TOOLS_VERSION))

.PHONY: setup-envtest
setup-envtest: envtest ## Download the binaries required for ENVTEST in the local bin directory.
	@echo "Setting up envtest binaries for Kubernetes version $(ENVTEST_K8S_VERSION)..."
	@$(ENVTEST) use $(ENVTEST_K8S_VERSION) --bin-dir $(LOCALBIN) -p path || { \
		echo "Error: Failed to set up envtest binaries for version $(ENVTEST_K8S_VERSION)."; \
		exit 1; \
	}

.PHONY: envtest
envtest: $(ENVTEST) ## Download setup-envtest locally if necessary
$(ENVTEST): $(LOCALBIN)
	$(call go-install-tool,$(ENVTEST),$(ENVTEST_BIN_URL),$(ENVTEST_VERSION))

.PHONY: golangci-lint
golangci-lint: $(GOLANGCI_LINT) ## Download golangci-lint locally if necessary
$(GOLANGCI_LINT): $(LOCALBIN)
	$(call go-install-tool,$(GOLANGCI_LINT),$(GOLANGCI_LINT_BIN_URL),$(GOLANGCI_LINT_VERSION))

# go-install-tool will 'go install' any package with custom target and name of binary, if it doesn't exist
# $1 - target path with name of binary
# $2 - package url which can be installed
# $3 - specific version of package
define go-install-tool
@[ -f "$(1)-$(3)" ] || { \
set -e; \
package=$(2)@$(3) ;\
echo "Downloading $${package}" ;\
rm -f $(1) || true ;\
GOBIN=$(LOCALBIN) go install $${package} ;\
mv $(1) $(1)-$(3) ;\
} ;\
ln -sf $(1)-$(3) $(1)
endef
