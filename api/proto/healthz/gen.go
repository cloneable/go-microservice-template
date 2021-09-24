//go:generate protoc -I. -I../../../vendor/github.com/googleapis/googleapis/ --go_out . --go_opt paths=source_relative healthz.proto
//go:generate protoc -I. -I../../../vendor/github.com/googleapis/googleapis/ --go-grpc_out . --go-grpc_opt paths=source_relative --go-grpc_opt=require_unimplemented_servers=false healthz.proto
//go:generate protoc -I. -I../../../vendor/github.com/googleapis/googleapis/ --grpc-gateway_out . --grpc-gateway_opt paths=source_relative healthz.proto
//go:generate protoc -I. -I../../../vendor/github.com/googleapis/googleapis/ --openapiv2_out . healthz.proto

package healthz
