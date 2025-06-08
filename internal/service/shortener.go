// shortener.go
package service

import "url_shortener/internal/storage"

type ShortenService struct {
	store storage.Store
}

func NewShortenService(s storage.Store) *ShortenService {
	return &ShortenService{store: s}
}

func (s *ShortenService) ShortenURL(originalURL string) string {
	return s.store.Save(originalURL)
}

func (s *ShortenService) ResolveCode(code string) (string, bool) {
	return s.store.Resolve(code)
}

func (s *ShortenService) GetTopDomains() map[string]int {
	return s.store.GetTopDomains(3)
}
