package repository

type CacheRepository interface {
	Insert(key string) error
	Exists(key string) bool
}
