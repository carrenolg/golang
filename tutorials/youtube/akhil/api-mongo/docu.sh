# 1. create new module
$ go mod init github.com/carrenolg/go-mongo
# 2. install httprouter (routes handler)
$ go get -u "github.com/julienschmidt/httprouter"
# 3. install driver mongo
$ go get -u "gopkg.in/mgo.v2"
# 4. install mongo bson
$ go get -u "gopkg.in/mgo.v2/bson"
# 5. run
$ go run main.go
# 6. auth app (golang)
# create docker container mongo
# "mongodb://root:admin@localhost:27017"
