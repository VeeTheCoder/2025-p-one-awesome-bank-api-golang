package handlers

import (
    "github-api-boilerplate/models"
    "github.com/gin-gonic/gin"
)

type Auth struct{}

func (a *Auth) Register(g *gin.RouterGroup) {
    g.POST("/register", a.register)
}

func (a *Auth) register(c *gin.Context) {
    var user models.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": "invalid request"})
        return
    }
    repo := repositories.User{}
    if err := repo.Create(user); err != nil {
        c.JSON(500, gin.H{"error": "internal server error"})
        return
    }
    c.JSON(201, gin.H{"message": "user created successfully"})
}