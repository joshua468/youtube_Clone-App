package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/youtube_Clone/models"
	"github.com/joshua468/youtube_Clone/services"
)

type UserHandlers struct {
	Service services.UserService 
}

func NewUserHandlers(service services.UserService) *UserHandlers {
	return &UserHandlers{
		Service: service,
	}
}

func (h *UserHandlers) RegisterRoutes(router *gin.Engine) {
	router.POST("/users", h.CreateUser)
	router.GET("/users/:id", h.GetUserByID)
	router.PUT("/users/:id", h.UpdateUser)
	router.DELETE("/users/:id", h.DeleteUser)
	router.GET("/users", h.ListUsers)
}

func (h *UserHandlers) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}


func (h *UserHandlers) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := h.Service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandlers) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.UpdateUser(id, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}


func (h *UserHandlers) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := h.Service.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *UserHandlers) ListUsers(c *gin.Context) {
	users, err := h.Service.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
