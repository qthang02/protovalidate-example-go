### đảm bảo có protoc-gen trong máy: 
    `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
### cài thêm các  package 

`go get github.com/envoyproxy/protoc-gen-validate`
`go get github.com/bufbuild/protovalidate-go`

lý do:
`protovalidate` là project phát triển từ `protoc-gen-validate`, và do đó có một số phụ thuộc vào `protoc-gen-validate`


