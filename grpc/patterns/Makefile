gen-proto:
	protoc proto/*.proto \
		--proto_path=./proto \
		--go_out=./proto --go_opt=paths=source_relative \
		--go-grpc_out=require_unimplemented_servers=false:./proto --go-grpc_opt=paths=source_relative