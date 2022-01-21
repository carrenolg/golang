# 1. install protoc on ubuntu
$curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.19.3/protoc-3.19.3-linux-x86_64.zip
$ unzip protoc-3.19.3-linux-x86_64.zip -d protoc3
$ sudo mv protoc3/bin/protoc /usr/bin/protoc
# 2. install package (from go.mod directory)
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# 3. Set up directory and variable enviroment
SRC_DIR=/home/gio10/cs/dev/golang/books/gorest/src/ch6/01-proto
DST_DIR=/home/gio10/cs/dev/golang/books/gorest/src/ch6/01-proto/out
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/person.proto
