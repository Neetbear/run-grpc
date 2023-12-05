# run-grpc
study grpc

## install
``` cmd
$ brew install protobuf
$ go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
$ which protoc-gen-go-grpc
```

## proto buf pb
```cmd
$ protoc --go_out=. --go-grpc_out=. proto/greet.proto
$ go mod tidy
```