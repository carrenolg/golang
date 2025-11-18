# Set up mysql container
docker run -d --name microservice-api-mysql -e TZ=UTC -p 30306:3306 -e MYSQL_ROOT_PASSWORD=1q2w3e4r -e MYSQL_USER=admin -e MYSQL_DATABASE=banking -e MYSQL_PASSWORD=password ubuntu/mysql:8.0-22.04_beta
