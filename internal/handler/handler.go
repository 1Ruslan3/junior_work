package handler

import (
	"net/http"
	"strconv"
	"time"

	"junior/internal/model"
	"junior/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.SubscriptionService
}

func NewHandler(s *service.SubscriptionService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.POST("/subscriptions", h.Create)
	r.GET("/subscriptions", h.GetAll)
	r.GET("/subscriptions/:id", h.GetByID)
	r.PUT("/subscriptions/:id", h.Update)
	r.DELETE("/subscriptions/:id", h.Delete)
	r.GET("/subscriptions/total", h.CalculateTotal)
}

func (h *Handler) Create(c *gin.Context) {
	var sub model.Subscription
	if err := c.ShouldBindJSON(&sub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(&sub); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, sub)
}

func (h *Handler) GetAll(c *gin.Context) {
	subs, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subs)
}

func (h *Handler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sub, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, sub)
}

func (h *Handler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var sub model.Subscription
	if err := c.ShouldBindJSON(&sub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sub.ID = uint(id)
	if err := h.service.Update(&sub); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sub)
}

func (h *Handler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) CalculateTotal(c *gin.Context) {
	fromStr := c.Query("from")
	toStr := c.Query("to")
	userID := c.Query("user_id")
	serviceName := c.Query("service_name")

	from, err := time.Parse("2006-01", fromStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid from date"})
		return
	}
	to, err := time.Parse("2006-01", toStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid to date"})
		return
	}
	total, err := h.service.CalculateTotal(from, to, userID, serviceName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": total})
}
