.PHONY: generate-go e2e-go

generate-go:
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen \
		--config oapi-codegen.yaml ../openapi.yaml

e2e-go:
	go test ./client/ -count=1 -v
