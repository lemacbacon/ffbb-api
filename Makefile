.PHONY: lint lint-yamls lint-openapi lint-workflows generate e2e

include go/Makefile

lint: lint-yamls lint-openapi lint-workflows

lint-yamls:
	yamllint -d "{extends: relaxed, rules: {line-length: {max: 120}}}" .

lint-openapi:
	npx --yes @redocly/cli lint openapi.yaml

lint-workflows:
	go run github.com/rhysd/actionlint/cmd/actionlint@latest

generate: generate-go

e2e: e2e-go
