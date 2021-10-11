TO RUN DOCKERIZED TODO LIST APPLICATION

```
docker-compose build
docker-compose up
```

------------------------------

> Docker compose builds and starts db, backend web api and frontend applications. If you want to up applications seperately, run commands below.

TO CREATE DATABASE ONLY, RUN FOLLOWING COMMANDS:

```
docker run -d -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=root mysql
docker exec -it mysql mysql -uroot -proot -e 'CREATE DATABASE tododb CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci'
```

TO START BACKEND APPLICATION ONLY, RUN FOLLOWING COMMANDS AFTER DATABASE CREATION:

```
cd server
go run .
```


TO START FRONTEND APPLICATION RUN FOLLOWING COMMANDS:

```
cd client
ng serve -o
```


TO START TESTING OF FRONTEND APPLICATION RUN FOLLOWING COMMANDS:

```
cd client
ng test
```

TO START TESTING OF BACKEND APPLICATION RUN FOLLOWING COMMANDS:

```
cd server
go test ./...
```
