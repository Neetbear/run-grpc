# run-grpc
study grpc

## protobuf install
``` console
$ brew install protobuf
$ go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
$ which protoc-gen-go-grpc
```

## proto buf pb
``` console
$ protoc --go_out=. --go-grpc_out=. proto/greet.proto
$ go mod tidy
```
``` console
$ protoc --proto_path=proto proto/*.proto --go_out=gen/
$ protoc --proto_path=proto proto/*.proto --go-grpc_out=gen/
```

## go work
``` console
$ go work init
$ go work use ./tools/ ./tools/gopls/
```

## google api
``` console
$ mkdir -p google/pai
$ curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > google/api/annotations.proto
$ curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > google/api/http.proto
```
``` console
$ protoc -I . --grpc-gateway_out ./gen/ \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    proto/greet.proto
```

## mux
mux port != grpc port

## buf cli
``` console
$ brew install bufbuild/buf/buf
$ buf --help
```

## grpc-gateway
Proxy tool that converts Protocol Buffers and gRPC services into HTTP JSON endpoints.

``` console
$ go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## buf gen
``` yaml
version: v1
plugins:
  - name: go
    out: ../pkg
    opt:
      - paths=source_relative
  - name: go-grpc
    out: ../pkg
    opt:
      - paths=source_relative
  - name: grpc-gateway
    out: ../pkg
    opt:
      - paths=source_relative
      - generate_unbound_methods=true
  - name: openapiv2
    out: ../pkg
```

## proto gen validate
``` cmd
$ go get github.com/envoyproxy/protoc-gen-validate
```

if vscode ext error
```
"protoc": {
    "path": "/usr/local/bin/protoc",
    "options": [
        "--proto_path=${workspaceRoot}/common/proto", # -> check
    ]
}
```