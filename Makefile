run:
	@go run main.go

run-bloody:
	@go run services/bloody/*.go

run-sakura:
	@go run services/sakura/*.go
	
gen:
	@protoc \
		--proto_path=protobuf "protobuf/data.proto" \
		--go_out=services/common/genproto/crawler --go_opt=paths=source_relative \
  	--go-grpc_out=services/common/genproto/crawler --go-grpc_opt=paths=source_relative
