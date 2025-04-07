package main

import (
    "fmt"
    "log"

    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
    "github.com/swaggo/gin-swagger/swaggerFiles"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, "/swagger-ui/"))
    r.POST("/users", createUser)
    r.GET("/users/:id", getUser)
    r.Run(":8080")
}

func createUser(c *gin.Context) {
    var req CreateRequest
    if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }
    id := create(req)
    c.JSON(http.StatusCreated, CreateUserResponse{Id: id})
}

func getUser(c *gin.Context) {
    id := c.Param("id")
    user, err := getUser(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
        return
    }
    c.JSON(http.StatusOK, user)
}

func create(req CreateRequest) (int64, error) {
    // implement business logic here...
    return 1, nil
}

func getUser(id string) (*User, error) {
    // implement business logic here...
    return &User{}, nil
}