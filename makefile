gen:
	protoc --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:./pb --go-grpc_opt=paths=source_relative proto/*.proto

clean:
	rm -r pb/*

server:
	go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go -address 0.0.0.0:8080

