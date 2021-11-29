# 1. run stack.yml
$ docker-compose -f "stack.yml" up -d

# 2. come in to container
$ docker exec -it mysqldb /bin/bash
$ mysql -u root -p
# 3. create user into mysql server
# 172.19.0.1 is the address from your computrer assigned by docker bridge
mysql> CREATE USER 'dev'@'172.19.0.1' IDENTIFIED BY 'admin';
mysql> GRANT ALL PRIVILEGES ON * . * TO 'dev'@'172.19.0.1';
mysql> CREATE DATABASE web;
mysql> USE web;

# 5. create table
# db.sql script

# 6. run app
$ go run main.go