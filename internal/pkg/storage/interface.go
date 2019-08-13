package storage

type Storage interface {
	Add(key string, value string)
	Get(key string) []string
}
