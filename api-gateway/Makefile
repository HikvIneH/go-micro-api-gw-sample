proto:
	protoc --go_out=. --go-grpc_out=. pkg/auth/pb/*.proto
	protoc --go_out=. --go-grpc_out=. pkg/order/pb/*.proto
	protoc --go_out=. --go-grpc_out=. pkg/product/pb/*.proto

server:
	go run cmd/main.go