package main

import (
	"github.com/GugahBrz/go-api/conn"
	"github.com/GugahBrz/go-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors())

	v1 := r.Group("api/v1")
	{
		v1.POST("/users", postUser)
		v1.GET("/users", getUsers)
		v1.GET("/users/:id", getUser)
		v1.PUT("/users/:id", updateUser)
		v1.DELETE("/users/:id", deleteUser)
	}

	r.Run(":3001")
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func postUser(c *gin.Context) {

	db := conn.InitDb()
	defer db.Close()

	var user models.User
	c.Bind(&user)

	if user.Name != "" {
		db.Create(&user)
		c.JSON(201, gin.H{"success": user})
	} else {
		c.JSON(422, gin.H{"error": "Empty Fields"})
	}
}

func getUser(c *gin.Context) {

	db := conn.InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var user models.User
	db.First(&user, id)

	if user.ID != 0 {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

func getUsers(c *gin.Context) {

	db := conn.InitDb()
	defer db.Close()

	var users []models.User
	db.Find(&users)

	c.JSON(200, users)
}

func updateUser(c *gin.Context) {

	db := conn.InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var user models.User
	db.First(&user, id)

	if user.ID != 0 {
		var newUser models.User
		c.Bind(&newUser)

		if newUser.Name != "" {
			result := models.User{Name: newUser.Name}
			db.Save(&result)
			c.JSON(200, gin.H{"success": result})

		} else {
			c.JSON(422, gin.H{"error": "Empty Fields"})
		}
	} else {
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

func deleteUser(c *gin.Context) {

	db := conn.InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var user models.User
	db.First(&user, id)

	if user.ID != 0 {
		db.Delete(&user)
		c.JSON(200, gin.H{"success": "User " + id + " deleted"})
	} else {
		c.JSON(404, gin.H{"error": "User not found"})
	}
}
