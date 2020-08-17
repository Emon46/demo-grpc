# demo-grpc
simple project   to learn the use of  grpc (google Remote Procedure Calls)

#first install Protocol Buffer Compiler version-3
[install](https://grpc.io/docs/protoc-installation/)

# to generate  *.pb.go from .proto
```
$ protoc \
  --go_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:. \
  --go-grpc_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:. \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
   helloToDemo/hello.proto
```
