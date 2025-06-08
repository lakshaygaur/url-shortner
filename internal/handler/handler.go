package handler

import (
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"sort"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

// Simple in-memory store
var (
	urlStore      = make(map[string]string) // shortCode -> originalURL
	reverseLookup = make(map[string]string) // originalURL -> shortCode
	domainCount   = make(map[string]int)    // domain -> count
	storeLock     = sync.RWMutex{}
)

type shortenRequest struct {
	URL string `json:"url" binding:"required"`
}

type shortenResponse struct {
	ShortURL string `json:"short_url"`
}

// RegisterRoutes sets up all API routes
func RegisterRoutes(r *gin.Engine) {
	r.POST("/shorten", shortenURLHandler)
	r.GET("/r/:code", redirectHandler)
	r.GET("/metrics/domains", domainMetricsHandler)
}

func shortenURLHandler(c *gin.Context) {
	var req shortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	storeLock.RLock()
	if code, exists := reverseLookup[req.URL]; exists {
		storeLock.RUnlock()
		c.JSON(http.StatusOK, shortenResponse{ShortURL: "/r/" + code})
		return
	}
	storeLock.RUnlock()

	// Generate short code
	hash := sha1.Sum([]byte(req.URL))
	code := base64.URLEncoding.EncodeToString(hash[:])[:6]

	storeLock.Lock()
	urlStore[code] = req.URL
	reverseLookup[req.URL] = code
	domain := extractDomain(req.URL)
	domainCount[domain]++
	storeLock.Unlock()

	c.JSON(http.StatusOK, shortenResponse{ShortURL: "/r/" + code})
}

func redirectHandler(c *gin.Context) {
	code := c.Param("code")

	storeLock.RLock()
	originalURL, exists := urlStore[code]
	storeLock.RUnlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}

	c.Redirect(http.StatusFound, originalURL)
}

func domainMetricsHandler(c *gin.Context) {
	storeLock.RLock()
	type kv struct {
		Domain string
		Count  int
	}
	var list []kv
	for domain, count := range domainCount {
		list = append(list, kv{domain, count})
	}
	storeLock.RUnlock()

	sort.Slice(list, func(i, j int) bool {
		return list[i].Count > list[j].Count
	})

	top := make(map[string]int)
	for i := 0; i < len(list) && i < 3; i++ {
		top[list[i].Domain] = list[i].Count
	}

	c.JSON(http.StatusOK, top)
}

func extractDomain(url string) string {
	// remove scheme
	if strings.HasPrefix(url, "http://") {
		url = strings.TrimPrefix(url, "http://")
	} else if strings.HasPrefix(url, "https://") {
		url = strings.TrimPrefix(url, "https://")
	}
	// split and take domain part
	parts := strings.Split(url, "/")
	return parts[0]
}
