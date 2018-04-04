package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.POST("/users", PostUser)
		v1.GET("/users", GetUsers)
		v1.GET("/users/:id", GetUser)
		v1.PUT("/users/:id", UpdateUser)
		v1.DELETE("/users/:id", DeleteUser)
	}

	r.Run(":8080")
}

func PostUser(c *gin.Context) {

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"name\": \"Zanoni\" }" http://localhost:8080/api/v1/users}
}

func GetUser(c *gin.Context) {

	// curl -i http://localhost:8080/api/v1/users/1
}

func GetUsers(c *gin.Context) {

	// curl -i http://localhost:8080/api/v1/users
}

func UpdateUser(c *gin.Context) {

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \name\": \"Zanoni\" }" http://localhost:8080/api/v1/users/1}
}

func DeleteUser(c *gin.Context) {

	// curl -i -X DELETE http://localhost:8080/api/v1/users/1
}
