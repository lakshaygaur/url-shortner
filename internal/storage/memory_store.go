// memory_store.go
package storage

import (
	"crypto/sha1"
	"encoding/base64"
	"sort"
	"strings"
	"sync"
)

type MemoryStore struct {
	urlToCode map[string]string
	codeToURL map[string]string
	domainUse map[string]int
	mu        sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		urlToCode: make(map[string]string),
		codeToURL: make(map[string]string),
		domainUse: make(map[string]int),
	}
}

func (s *MemoryStore) Save(url string) (string, bool) {
	hash := sha1.Sum([]byte(url))
	code := base64.URLEncoding.EncodeToString(hash[:])[:6]

	s.mu.Lock()
	defer s.mu.Unlock()
	s.urlToCode[url] = code
	s.codeToURL[code] = url
	domain := extractDomain(url)
	s.domainUse[domain]++

	return code, false
}

func (s *MemoryStore) Resolve(code string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	url, ok := s.codeToURL[code]
	return url, ok
}

func (s *MemoryStore) GetTopDomains(limit int) map[string]int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	type kv struct {
		Key   string
		Value int
	}
	var list []kv
	for domain, count := range s.domainUse {
		list = append(list, kv{domain, count})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Value > list[j].Value
	})

	top := make(map[string]int)
	for i := 0; i < len(list) && i < limit; i++ {
		top[list[i].Key] = list[i].Value
	}
	return top
}

func extractDomain(url string) string {
	url = strings.TrimPrefix(url, "https://")
	url = strings.TrimPrefix(url, "http://")
	return strings.Split(url, "/")[0]
}
