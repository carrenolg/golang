# Section 3 - Simple Testing
$ go test -v # run all tests
# coverage report
$ go test -coverprofile=coverage.out
$ go tool cover -html=coverage.out
