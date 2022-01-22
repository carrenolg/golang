# Due The Builder 2019
1. Create mcache.proto
2. protoc --go_out=. --go-grpc_out=. --plugin=grpc mcache.proto
3. go get google.golang.org/protobuf/proto # run from go.mod directory
4. go get google.golang.org/grpc # run from go.mod directory