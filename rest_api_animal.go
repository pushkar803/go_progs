package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Animal struct {
	id   string
	name string
}

var animals = []Animal{
	{
		id:   "1",
		name: "one",
	},
	{
		id:   "2",
		name: "two",
	},
}

func main9() {

	r := gin.Default()

	authGrp := r.Group("/auth")
	{
		authGrp.Use(authMiddleware)
		authGrp.GET("/animal", GetAllAnimals)
		authGrp.GET("/animal/:id", GetAllAnimalByID)
		authGrp.POST("/animal", CreateAnimal)
		authGrp.PUT("/animal/:id", UpdateAnimalByID)
		authGrp.DELETE("/animal/:id", DeleteAnimalById)
	}

	r.Run(":8085")
}

func authMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		c.Abort()
		return
	}

	token := authHeader[len("Bearer "):]
	if !validToken(token) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		c.Abort()
		return
	}
	c.Next()
}

func validToken(token string) bool {
	return true
}

func GetAllAnimals(c *gin.Context) {
	c.JSON(http.StatusOK, animals)
}

func GetAllAnimalByID(c *gin.Context) {
	id := c.Param("id")
	for _, animal := range animals {
		if animal.id == id {
			c.JSON(http.StatusOK, animal)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
}

func CreateAnimal(c *gin.Context) {

	var newAnimal Animal
	err := c.ShouldBindJSON(newAnimal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	}

	animals = append(animals, newAnimal)
	c.JSON(http.StatusOK, gin.H{"message": "new animal created"})
}

func UpdateAnimalByID(c *gin.Context) {
	id := c.Param("id")
	var updatedAnimal Animal
	err := c.ShouldBindJSON(updatedAnimal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	}
	for i, animal := range animals {
		if animal.id == id {
			animals[i] = updatedAnimal
			c.JSON(http.StatusOK, gin.H{"message": "animal updated"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
}

func DeleteAnimalById(c *gin.Context) {

	id := c.Param("id")
	for i, animal := range animals {
		if animal.id == id {
			animals = append(animals[:i], animals[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "deleted success"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
}
