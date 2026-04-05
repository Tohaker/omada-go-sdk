fix:
	go run tools/fix_spec.go

generate:
	docker run --rm \
		-w /local \
		-v "${PWD}:/local" \
		"openapitools/openapi-generator-cli:v7.21.0" \
		batch generator/batch.yaml

tidy:
	go mod tidy

verify:
	go build ./...
	go vet ./...

all: fix generate tidy verify