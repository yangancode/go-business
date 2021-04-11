GOPATH:=$(shell go env GOPATH)


.PHONY: proto
greeter_go:
	protoc --proto_path=${GOPATH}/src:.  --go_out=plugins=grpc:.  proto/greeter/greeter.proto

.PHONY: proto
greeter_js:
	protoc --proto_path=${GOPATH}/src:.  --js_out=import_style=commonjs,binary:. proto/greeter/greeter.proto

.PHONY: proto
route_go:
	protoc --proto_path=${GOPATH}/src:.  --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/route/route.proto

.PHONY: proto
practice_go:
	protoc --proto_path=${GOPATH}/src:.  --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/practice/practice.proto