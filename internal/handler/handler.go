package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up all API routes
func RegisterRoutes(r *gin.Engine) {
	r.POST("/shorten", shortenURLHandler)
	r.GET("/r/:code", redirectHandler)
	r.GET("/metrics/domains", domainMetricsHandler)
}

func shortenURLHandler(c *gin.Context) {
	c.String(http.StatusOK, "Shorten URL endpoint - To be implemented")
}

func redirectHandler(c *gin.Context) {
	code := c.Param("code")
	c.String(http.StatusOK, "Redirect for code: %s - To be implemented", code)
}

func domainMetricsHandler(c *gin.Context) {
	c.String(http.StatusOK, "Domain metrics endpoint - To be implemented")
}
