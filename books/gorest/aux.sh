# install new package
$ go get github.com/justinas/alice
# keep in mind: you must run command from the /gorest folder
# Library list
go get github.com/gorilla/handlers
go get github.com/gorilla/rpc
go get github.com/emicklei/go-restful/v3
go get github.com/mattn/go-sqlite3
go get github.com/revel/revel
go get gopkg.in/mgo.v2
go get gopkg.in/mgo.v2/bson
# Debugging on Go
1. Create .vscode/launch.json
2. install dev, run following command from go.mod directory
go install github.com/go-delve/delve/cmd/dlv@latest
3. click on "install all" when pop-up is showed
# Insatll grpc
go get google.golang.org/grpc
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get github.com/lib/pq