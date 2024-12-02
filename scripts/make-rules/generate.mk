# ==============================================================================
# Makefile helper functions for generate necessary files
#


SERVICES ?= $(filter-out tools,$(foreach service,$(wildcard ${ONEX_ROOT}/cmd/*),$(notdir ${service})))

.PHONY: gen.run
gen.run: gen.docgo.doc gen.protoc go.generate ## Generate all necessary files except for files generated by `generated.files`.

# Notice: gen.apisprotobuf must after generated.files, or gen.apisprotobuf will fail
.PHONY: gen.k8s.all
gen.k8s.all: gen.docs generated.files gen.apisprotobuf # Generate all necessary files related to kubernetes.

.PHONY: gen.all
gen.all: gen.run gen.k8s.all gen.appdocs gen.helm-docs ## Generate all possible project code that can be generated.

.PHONY: gen.docgo.doc
gen.docgo.doc: ## Generate missing doc.go for go packages.
	@echo "===========> Generating missing doc.go for go packages"
	@${SCRIPTS_DIR}/gen-doc.sh

.PHONY: gen.docgo.check
gen.docgo.check: gen.docgo.doc ## Check missing doc.go for go packages.
	@n="$$(git ls-files --others '*/doc.go' | wc -l)"; \
	if test "$$n" -gt 0; then \
		git ls-files --others '*/doc.go' | sed -e 's/^/  /'; \
		echo "$@: untracked doc.go file(s) exist in working directory" >&2 ; \
		false ; \
	fi

.PHONY: gen.docs
gen.docs: ## Update generated swagger docs.
	@${SCRIPTS_DIR}/update-codegen.sh swagger

.PHONY: gen.appdocs
gen.appdocs: ## Update generated application docs.
	@${SCRIPTS_DIR}/update-generated-docs.sh

.PHONY: gen.helm-docs
gen.helm-docs: tools.verify.helm-docs ## Generate documentation for helm charts.
	@echo "===========> Generating documentation for helm charts"
	@helm-docs --log-level=warning --chart-search-root $(ONEX_ROOT)

.PHONY: gen.ca.%
gen.ca.%: ## Generate CA files.
	$(eval CA := $(word 1,$(subst ., ,$*)))
	@echo "===========> Generating CA files for $(CA)"
	@${SCRIPTS_DIR}/gen-certs.sh generate-node-cert $(OUTPUT_DIR)/cert $(CA)

.PHONY: gen.ca
gen.ca: $(addprefix gen.ca., $(CERTIFICATES)) ## Generate all CA files.

.PHONY: gen.kubeconfig
gen.kubeconfig: gen.ca ## Generate kubeconfig files.
	@echo "===========> Generating admin kubeconfig file"
	@$(SCRIPTS_DIR)/gen-kubeconfig.sh $(OUTPUT_DIR)/cert/ca.pem \
		$(OUTPUT_DIR)/cert/admin.pem $(OUTPUT_DIR)/cert/admin-key.pem > \
		$(OUTPUT_DIR)/config

.PHONY: gen.protoc
gen.protoc: ## Generate go source files from protobuf files.
	@protoc \
		--proto_path=$(APIROOT) \
		--proto_path=$(APISROOT) \
		--proto_path=$(ONEX_ROOT)/third_party/protobuf \
		--go_out=paths=source_relative:$(APIROOT) \
		--go-http_out=paths=source_relative:$(APIROOT) \
		--go-grpc_out=paths=source_relative:$(APIROOT) \
		--go-errors_out=paths=source_relative:$(APIROOT) \
		--go-errors-code_out=paths=source_relative:$(ONEX_ROOT)/docs/guide/zh-CN/api/errors-code \
		--validate_out=paths=source_relative,lang=go:$(APIROOT) \
		--openapi_out=fq_schema_naming=true,default_response=false:$(ONEX_ROOT)/api/openapi \
		--openapiv2_out=$(ONEX_ROOT)/api/openapi \
		--openapiv2_opt=logtostderr=true \
		--openapiv2_opt=json_names_for_fields=false \
		$(shell find $(APIROOT) -name *.proto)
	# Only onex-fake-server use grpc-gateway
	@protoc \
		--proto_path=$(APIROOT) \
		--proto_path=$(APISROOT) \
		--proto_path=$(ONEX_ROOT)/third_party/protobuf \
		--grpc-gateway_out=paths=source_relative:$(APIROOT) \
		$(shell find $(APIROOT)/fakeserver -name *.proto)

.PHONY: gen.apisprotobuf
gen.apisprotobuf: ## Generate protobuf files from structs.
	#@$(SCRIPTS_DIR)/update-codegen.sh
	@cp $(ONEX_ROOT)/manifests/generated.pb.go.fix $(ONEX_ROOT)/pkg/apis/apps/v1beta1/generated.pb.go

.PHONY: go.generate
go.generate: tools.verify.mockgen tools.verify.wire ## Run `go generate ./...` command.
	@$(GO) generate $(ONEX_ROOT)/...

.PHONY: gen.systemd
gen.systemd: $(addprefix gen.systemd., $(SERVICES)) ## Generate all systemd unit files.

.PHONY: gen.systemd.%
gen.systemd.%: ## Generate specified systemd unit file.
	$(eval ONEX_ENV_FILE ?= $(MANIFESTS_DIR)/env.local)
	$(eval GENERATED_SERVICE_DIR := $(OUTPUT_DIR)/systemd)
	$(eval SERVICE := $(lastword $(subst ., ,$*)))
	@echo "===========> Generating $(SERVICE) systemd unit file"
	@$(SCRIPTS_DIR)/gen-service-config.sh $(SERVICE) $(ONEX_ENV_FILE) \
		$(ONEX_ROOT)/configs/onex.systemd.tmpl.service $(GENERATED_SERVICE_DIR)
ifeq ($(V),1)
	echo "DBG: Generating systemd unit file at $(GENERATED_SERVICE_DIR)/$(SERVICE).service"
endif

.PHONY: gen.appconfig
gen.appconfig: $(addprefix gen.appconfig., $(SERVICES)) ## Generate all application configuration files.

.PHONY: gen.appconfig.%
gen.appconfig.%: ## Generate specified application configuration file.
	$(eval ONEX_ENV_FILE ?= $(MANIFESTS_DIR)/env.local)
	$(eval GENERATED_SERVICE_DIR := $(OUTPUT_DIR)/appconfig)
	$(eval SERVICE := $(lastword $(subst ., ,$*)))
	@echo "===========> Generating $(SERVICE) configuration file"
	@$(SCRIPTS_DIR)/gen-service-config.sh $(SERVICE) $(ONEX_ENV_FILE) \
		$(ONEX_ROOT)/configs/appconfig/$(SERVICE).config.tmpl.yaml $(GENERATED_SERVICE_DIR)
ifeq ($(V),1)
	echo "DBG: Generating $(SERVICE) application configuration file at $(GENERATED_SERVICE_DIR)/$(SERVICE)"
endif