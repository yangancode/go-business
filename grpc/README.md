# gRPC Demo


## Download
### protobuf
```shell script
$ export GO111MODULE=on  # Enable module mode
$ go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc
```


## Error
1.undefined: grpc.SupportPackageIsVersion7
```
// error
# go-business/grpc/proto/route
proto/route/route_grpc.pb.go:15:11: undefined: grpc.SupportPackageIsVersion7
proto/route/route_grpc.pb.go:199:33: undefined: grpc.ServiceRegistrar

// solution
grpc version must greater than  v1.32.0
```