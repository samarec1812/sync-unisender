.PHONY: build
build:
	$(info Build unisender-api images...)
	docker build -t unisender-api .


.PHONY: lint
lint:
	$(info Run go linters in project...)
	golangci-lint run ./... -c ./configs/.golangci.yml