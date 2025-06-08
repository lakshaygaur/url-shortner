package storage

type Store interface {
	Save(url string) string
	Resolve(code string) (string, bool)
	GetTopDomains(limit int) map[string]int
	ClearStorage()
}
