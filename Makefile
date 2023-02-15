export GOOSE_DRIVER?=mysql
export GOOSE_DBSTRING?=root:qwerty@tcp(localhost:3307)/amo_auth?charset=utf8

.PHONY: proto
PROTOC_GEN_GO := $(GOPATH)/bin/protoc-gen-go

# If $GOPATH/bin/protoc-gen-go does not exist, we'll run this command to install it.
$(PROTOC_GEN_GO):
	go get -u github.com/golang/protobuf/protoc-gen-go

proto:
	$(info Generate proto...)
	protoc -I ./proto \
      --go_out ./proto/auth --go_opt paths=source_relative \
      --go-grpc_out ./proto/auth --go-grpc_opt paths=source_relative \
      --grpc-gateway_out ./proto/auth --grpc-gateway_opt paths=source_relative \
      --openapiv2_out ./api --openapiv2_opt logtostderr=true \
      ./proto/auth.proto


.PHONY: lint
lint:
	$(info Run go linters in project...)
	golangci-lint run ./.../... -c ./.golangci.yml

.PHONY: build-auth
build-auth:
	$(info Build unicender-api images...)
	docker build -f docker/auth/Dockerfile -t auth-api .


.PHONY: build-unisender
build-unisender:
	$(info Build unisender-api images...)
	docker build -f docker/unisender/Dockerfile -t unisender-api .


.PHONY: up
up:
	$(info Run services...)
	docker-compose up --build


.PHONY: down
down:
	$(info Stop services and remove containers...)
	docker-compose down


.PHONY: migrate-up
migrate-up:
	$(info Execute migrations in project...)
	goose -dir migrations up

.PHONY: migrate-down
migrate-down:
	$(info Rollback migrations in project...)
	goose -dir migrations down


.PHONY: db-status
db-status:
	$(info Print status database in project...)
	goose status

.PHONY: swagger
swagger:
	$(info Generate swagger documentation in project...)
	swagger generate spec -o ./api/unisender.swagger.json --scan-models


.PHONY: api-doc
api-doc:
	$(info Serve documentation on localhost...)
	swagger serve -F unisender.swagger ./api/unisender.swagger.json auth.swagger ./api/auth.swagger.json

.PHONY: swag-markdown
swag-markdown:
	$(info Generate markdown swagger documentation in project ...)
	swagger generate markdown -f ./api/unisender.swagger.json --output ./api/unisender.swagger.md
	swagger generate markdown -f ./api/auth.swagger.json --output ./api/auth.swagger.md