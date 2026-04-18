fix:
	go run tools/fix_spec.go

clean:
	out="omada"; \
    find "$$out" -mindepth 1 \
        \( -path "$$out/.gitignore" -o -path "$$out/.openapi-generator-ignore" -o -path "$$out/.nav.yml" \) -prune \
        -o -exec rm -rf {} +

generate: clean
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