package storage

type Store interface {
	Save(url string) (string, bool)
	Resolve(code string) (string, bool)
	GetTopDomains(limit int) map[string]int
}
