# go-api

An basic API to demonstrate Go, gin and gorm

# Schema

Table 'user'
id: integer / auto-increment
name: varchar(255) / Not-Null

Routes
POST : http://127.0.0.1:8080/api/v1/users
GET : http://127.0.0.1:8080/api/v1/users
GET : http://127.0.0.1:8080/api/users/1
PUT : http://127.0.0.1:8080/api/users/1
DELETE : http://127.0.0.1:8080/api/users/1

# Dependencies 

Gin for routes
go get github.com/gin-gonic/gin

Gorm for ORM
go get github.com/jinzhu/gorm