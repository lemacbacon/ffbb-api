.PHONY: lint lint-workflows generate generate-go verify-go e2e

lint:
	npx --yes @redocly/cli lint openapi.yaml

lint-workflows:
	go run github.com/rhysd/actionlint/cmd/actionlint@latest

generate: generate-go

generate-go:
	cd go && go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen \
		--config oapi-codegen.yaml ../openapi.yaml

verify-go:
	$(MAKE) generate-go
	git diff --exit-code go/client/client.gen.go

e2e:
	cd go && go test ./client/ -count=1 -v
