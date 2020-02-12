.PHONY: dep
dep:
	go mod tidy
	go mod vendor

.PHONY: lint
lint:
	export GOFLAGS=-mod=vendor
	golangci-lint run

.PHONY: test
test:
	go test -mod=vendor -gcflags=all=-l $(shell go list ./... | grep -v examples) -covermode=count -coverprofile .coverage.cov
	go tool cover -func=.coverage.cov
