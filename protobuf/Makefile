GOPATH:=$(shell go env GOPATH)

# proto1 相关指令
.PHONY: proto1
proto1:
	protoc --go_out=. proto1/greeter/greeter.proto
proto2:
	protoc -I. --go_out=. proto1/greeter/greeter.proto # 点号表示当前路径，注意-I参数没有等于号
proto3:
	protoc --proto_path=. --go_out=. proto1/greeter/greeter.proto
# 报错，找不到文件
proto4:
	protoc --proto_path=./proto/user --go_out=. proto1/greeter/greeter.proto

# 保存到其他目录，不使用该目录的包名
proto5:
	protoc --proto_path=. --go_out=./proto1/pb_go proto1/greeter/greeter.proto
# 保存到其他目录，使用该目录的包名
proto6:
	protoc --proto_path=. --go_out=. proto1/greeter/greeter_v2.proto

.PHONY: path
path1:
	protoc --proto_path=. --go_out=. proto1/greeter/greeter_v2.proto
path2:
	protoc --proto_path=. --go_out=paths=import:. proto1/greeter/greeter_v2.proto
path3:
	protoc --proto_path=. --go_out=paths=source_relative:. proto1/greeter/greeter_v2.proto


.PHONY: grpc
plugin1:
	protoc --proto_path=. --go_out=plugins=grpc,paths=import:./proto1/grpc_go proto1/greeter/greeter_v2.proto


.PHONY: import
import1:
	protoc --proto_path=. --go_out=. proto2/greeter/greeter.proto proto2/common.proto


# 正确的写法
multi1:
	protoc --proto_path=.  --go_out=.  proto/*.proto proto/user/*proto proto/greeter/*proto
# 错误写法，只会编译 common.proto
multi2:
	protoc --proto_path=.  --go_out=.  proto/*.proto

delete_pb:
	rm -rf proto/common
	rm -rf proto/greeter/*.pb.go
	rm -rf proto/user/*.pb.go





