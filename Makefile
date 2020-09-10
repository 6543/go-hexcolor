.PHONY: clean
clean:
	go clean -i ./...

.PHONY: fmt
fmt:
	go fmt *.go

.PHONY: vet
vet:
	go vet github.com/6543/go-hexcolor

.PHONY: misspell
misspell:
	@hash misspell > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		GO111MODULE=off $(GO) get -u github.com/client9/misspell/cmd/misspell; \
	fi
	misspell -w *

.PHONY: test
test:
	go test

.PHONY: vendor
vendor:
	GO111MODULE=on; go mod tidy && go mod vendor
