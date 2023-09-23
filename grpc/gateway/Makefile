build:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/hello.proto
	protoc -I . --grpc-gateway_out . \
	--grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    proto/hello.proto

run-server:
	go run server/main.go

run-client:
	go run server/main.go

run-proxy:
	go run server/main.go