package test

import (
	"testing"

	"url_shortener/internal/service"
	"url_shortener/internal/storage"

	"github.com/stretchr/testify/assert"
)

func TestShortenService(t *testing.T) {
	store := storage.NewMemoryStore()
	svc := service.NewShortenService(store)

	// Test 1: Shorten a new URL
	code1 := svc.ShortenURL("https://example.com/page1")
	assert.NotEmpty(t, code1)

	// Test 2: Shorten the same URL should return same code
	code2 := svc.ShortenURL("https://example.com/page1")
	assert.Equal(t, code1, code2)

	// Test 3: Resolve code
	originalURL, ok := svc.ResolveCode(code1)
	assert.True(t, ok)
	assert.Equal(t, "https://example.com/page1", originalURL)

	// Test 4: Resolve invalid code
	_, ok = svc.ResolveCode("invalid")
	assert.False(t, ok)

	// Test 5: Metrics tracking
	store.ClearStorage()
	urls := []string{
		"https://youtube.com/a",
		"https://youtube.com/b",
		"https://youtube.com/c",
		"https://wikipedia.org/x",
		"https://udemy.com/y",
		"https://udemy.com/z",
		"https://udemy.com/k",
		"https://udemy.com/m",
	}
	for _, u := range urls {
		svc.ShortenURL(u)
	}
	top := svc.GetTopDomains()
	assert.Equal(t, 3, len(top))
	assert.Equal(t, 4, top["udemy.com"])
	assert.Equal(t, 3, top["youtube.com"])
	assert.Equal(t, 1, top["wikipedia.org"])
}
