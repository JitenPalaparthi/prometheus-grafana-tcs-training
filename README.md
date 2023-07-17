- Create and run postgres docker container

```
docker run -itd -e POSTGRES_USER=demo -e POSTGRES_PASSWORD=demo@123 -p 5432:5432 -e POSTGRES_DB=demodb --name postgresql postgres
```


docker run -itd -p 3306:3306 -e MYSQL_DATABASE=demodb -e MYSQL_USER=demo -e MYSQL_PASSWORD=demo@123  -e MYSQL_ROOT_PASSWORD=demo@123 --name mysqlc mysql


