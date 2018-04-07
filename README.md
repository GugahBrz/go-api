# go-api

An basic API to demonstrate Go, gin and gorm

### Schema

Table "User"
* id: integer / auto-increment
* name: varchar(255) / Not-Null

Routes
* POST : http://127.0.0.1:3001/api/v1/users
* GET : http://127.0.0.1:3001/api/v1/users
* GET : http://127.0.0.1:3001/api/users/1
* PUT : http://127.0.0.1:3001/api/users/1
* DELETE : http://127.0.0.1:3001/api/users/1

### Dependencies 

###### Gin for routes
```
go get github.com/gin-gonic/gin
```
###### Gorm for ORM
```
go get github.com/jinzhu/gorm
```
> You need to create a new schema in your database and configure him in this project.<br>
> You can configure the connection strings from conn/mysql.go

### Basic Usage

###### POST User
```
curl -i -X POST -H "Content-Type: application/json" -d "{ \"name\": \"Zanoni\" }" http://localhost:3001/api/v1/users}
```

###### GET User
```
curl -i http://localhost:3001/api/v1/users/1
```

###### GET Users
```
curl -i http://localhost:3001/api/v1/users
```

###### PUT User
```
curl -i -X PUT -H "Content-Type: application/json" -d "{ \name\": \"Zanoni\" }" http://localhost:3001/api/v1/users/1}
```

###### DELETE User
```
curl -i -X DELETE http://localhost:3001/api/v1/users/1
```
### Author
Gustavo Zanoni -
[LinkedIn](https://br.linkedin.com/in/gustavo-zanoni-6371a791 "LinkedIn Link")
