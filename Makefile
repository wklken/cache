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
	go test -mod=vendor -gcflags=all=-l $(shell go list ./...) -covermode=count -coverprofile .coverage.cov
