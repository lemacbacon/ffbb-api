.PHONY: lint lint-yamls lint-openapi lint-workflows generate generate-go e2e e2e-go

lint: lint-yamls lint-openapi lint-workflows

lint-yamls:
	yamllint -d "{extends: relaxed, rules: {line-length: {max: 120}}}" .

lint-openapi:
	npx --yes @redocly/cli lint openapi.yaml

lint-workflows:
	go run github.com/rhysd/actionlint/cmd/actionlint@latest

generate: generate-go

generate-go:
	$(MAKE) -C go generate-go

e2e: e2e-go

e2e-go:
	$(MAKE) -C go e2e-go
