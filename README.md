goctl api new demo
goctl rpc template -o demo.proto  
goctl rpc protoc demo.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.


go run demo.go -f etc/demo-api.yaml