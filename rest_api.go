package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	id   string
	name string
}

var users = []User{
	{id: "1", name: "one"},
	{id: "2", name: "two"},
}

func main8() {

	r := gin.Default()

	r.GET("/users", GetUsers)
	r.GET("/user/:id", GetUserById)
	r.POST("/users", CreateUser)
	r.PUT("/user/:id", UpdateUer)
	r.DELETE("/user/:id", DeleteUser)

	r.Run(":8084")
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	for _, user := range users {
		if user.id == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
}

func CreateUser(c *gin.Context) {

	var createUserPayload User
	err := c.ShouldBindJSON(createUserPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	lastIndex := len(users)
	users[lastIndex] = createUserPayload

	c.JSON(http.StatusOK, users)
}

func UpdateUer(c *gin.Context) {

	id := c.Param("id")
	var updateUserPayload User

	err := c.ShouldBindJSON(&updateUserPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	for i, user := range users {
		if user.id == id {
			users[i] = updateUserPayload
			c.JSON(http.StatusOK, gin.H{"message": "user updated"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "user deleted"})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	for i, user := range users {
		if user.id == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "user deleted"})
}
