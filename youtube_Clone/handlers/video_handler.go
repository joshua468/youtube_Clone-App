package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/youtube_Clone/models"
	"github.com/joshua468/youtube_Clone/services"
)

type VideoHandlers struct {
	Service services.VideoService
}

func NewVideoHandlers(service services.VideoService) *VideoHandlers {
	return &VideoHandlers{
		Service: service,
	}
}

func (h *VideoHandlers) RegisterRoutes(router *gin.Engine) {
	router.POST("/videos", h.CreateVideo)
	router.GET("/videos/:id", h.GetVideoByID)
	router.PUT("/videos/:id", h.UpdateVideo)
	router.DELETE("/videos/:id", h.DeleteVideo)
	router.GET("/videos", h.listVideos)
}

func (h *VideoHandlers) CreateVideo(c *gin.Context) {
	var video models.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateVideo(&video); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, video)
}

func (h *VideoHandlers) GetVideoByID(c *gin.Context) {
	id := c.Param("id")

	video, err := h.Service.GetVideoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, video)
}

func (h *VideoHandlers) UpdateVideo(c *gin.Context) {
	id := c.Param("id")

	var video models.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.UpdateVideo(id, &video); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, video)
}

func (h *VideoHandlers) DeleteVideo(c *gin.Context) {
	id := c.Param("id")

	if err := h.Service.DeleteVideo(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *VideoHandlers) ListVideos(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	videos, err := h.Service.ListVideos(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, videos)
}
