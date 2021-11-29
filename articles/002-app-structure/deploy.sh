# App is created for weight-tracker
# 1. Set up postgresql 
# run into 'resources' folder
$ docker-compose -f "stack.yml" up -d
# create new db and user
$ docker exec -it postgresdb /bin/bash
$ psql -U postgres # into container
>>postgre# create database 'weightracker'
>>postgre# create user 'dev'
>>postgre# alter user 'dev' with encrypted password 'admin';
>>postgre# grant all privileges on database 'weightracker' to 'dev';
# log in with dev user
$ psql -U dev -d weightracker # into container
