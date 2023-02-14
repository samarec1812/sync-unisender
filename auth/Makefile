.PHONY: proto
proto:
	$(info Generate proto...)
	protoc -I ./proto \
      --go_out ./proto --go_opt paths=source_relative \
      --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
      --grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
      --openapiv2_out ./proto  \
      ./proto/auth.proto


.PHONY: build
build:
	$(info Build unicender-api images...)
	docker build -t auth-api .
