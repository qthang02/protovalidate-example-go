build:
	protoc -I=protobuf --go_out=protobuf/ \
	  --go-grpc_out=protobuf/ \
      --validate_out="lang=go:protobuf" protobuf/greet.proto

