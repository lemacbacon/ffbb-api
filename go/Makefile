GO_DIR := go

.PHONY: generate-go e2e-go

generate-go:
	cd $(GO_DIR) && go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen \
		--config oapi-codegen.yaml ../openapi.yaml

e2e-go:
	cd $(GO_DIR) && go test ./client/ -count=1 -v
