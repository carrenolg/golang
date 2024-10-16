# check porcentage of coverage
go test --cover
# generate coverage file
go test -coverprofile=coverage.out
# generate html file
go tool cover -html=coverage.out -o coverage.html
# create users.sql
touch pkg/repository/dbrepo/testdata/users.sql
# setup colima with dockertest
sudo ln -s /Users/{user_name}/.colima/default/docker.sock /var/run/docker.sock
# keep in mind change your user name
# https://github.com/testcontainers/testcontainers-java/issues/5034