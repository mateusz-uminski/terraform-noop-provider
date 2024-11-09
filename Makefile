.PHONY: build
build:
	CGO_ENABLED=0 go build -o ./build/terraform-provider-noop

.PHONY: lint
lint:
	golangci-lint run -v

.PHONY: tests
tests:
	go test -v -race \
		-covermode atomic \
		-coverprofile coverage.out \
		./... -json > report.json
	go tool cover -func coverage.out

.PHONY: clean
clean:
	rm -rf build/
	rm -f coverage.out report.json
	rm -f examples/terraform.tfstate
	rm -f examples/terraform.tfstate.backup
