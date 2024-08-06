package controllers

import (
	"net/http"

	"github.com/balasl342/go-gin-starter-template/models"
	"github.com/balasl342/go-gin-starter-template/services"
	"github.com/balasl342/go-gin-starter-template/utils"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := services.GetUserById(id)
	if err != nil {
		utils.JSONResponse(c, http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	if err := services.CreateUser(&user); err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, gin.H{"message": "Could not create user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}
