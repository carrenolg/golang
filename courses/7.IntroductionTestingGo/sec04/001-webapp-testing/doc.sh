# check porcentage of coverage
go test --cover
# generate coverage file
go test -coverprofile=coverage.out
# generate html file
go tool cover -html=coverage.out -o coverage.html
