package handlers

import (
    "github-api-boilerplate/models"
    "github-api-boilerplate/repositories"
    "github.com/gin-gonic/gin"
)

type User struct{}

func (u *User) CreateUser(c *gin.Context) {
    var req models.CreateRequest
    if err := c.BindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": "invalid request"})
        return
    }
    repo := repositories.User{}
    user, err := repo.Create(req.User)
    if err != nil {
        c.JSON(500, gin.H{"error": "internal server error"})
        return
    }
    c.JSON(201, models.CreateUserResponse{Id: user.ID, User: user})
}

func (u *User) GetUsers(c *gin.Context) {
    repo := repositories.User{}
    users, err := repo.GetAll()
    if err != nil {
        c.JSON(500, gin.H{"error": "internal server error"})
        return
    }
    c.JSON(200, users)
}

func (u *User) UpdateUser(c *gin.Context) {
    var req models.CreateRequest
    if err := c.BindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": "invalid request"})
        return
    }
    repo := repositories.User{}
    user, err := repo.Update(req.User)
    if err != nil {
        c.JSON(500, gin.H{"error": "internal server error"})
        return
    }
    c.JSON(200, user)
}

func (u *User) DeleteUser(c *gin.Context) {
    var req models.CreateRequest
    if err := c.BindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": "invalid request"})
        return
    }
    repo := repositories.User{}
    err := repo.Delete(req.User)
    if err != nil {
        c.JSON(500, gin.H{"error": "internal server error"})
        return
    }
    c.JSON(200, gin.H{"message": "user deleted successfully"})
}