package handler

import (
	"net/http"

	"url_shortener/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	shortener *service.ShortenService
}

func NewHandler(svc *service.ShortenService) *Handler {
	return &Handler{shortener: svc}
}

type shortenRequest struct {
	URL string `json:"url" binding:"required"`
}

type shortenResponse struct {
	ShortURL string `json:"short_url"`
}

// RegisterRoutes sets up all API routes
func RegisterRoutes(r *gin.Engine, h *Handler) {
	r.POST("/shorten", h.shortenURLHandler)
	r.GET("/r/:code", h.redirectHandler)
	r.GET("/metrics/domains", h.domainMetricsHandler)
}

func (h *Handler) shortenURLHandler(c *gin.Context) {
	var req shortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	code := h.shortener.ShortenURL(req.URL)
	c.JSON(http.StatusOK, shortenResponse{ShortURL: "/r/" + code})
}

func (h *Handler) redirectHandler(c *gin.Context) {
	code := c.Param("code")
	originalURL, exists := h.shortener.ResolveCode(code)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}
	c.Redirect(http.StatusFound, originalURL)
}

func (h *Handler) domainMetricsHandler(c *gin.Context) {
	top := h.shortener.GetTopDomains()
	c.JSON(http.StatusOK, top)
}
