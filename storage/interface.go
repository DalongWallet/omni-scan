package storage

type Storage interface {
	Set(key string, value string) error

	Get(key string) (string, error)

	Delete(key string) error

	// scan key_start key_end limit
	// list key-value for (key_start, key_end]
	Scan(keyStart, keyEnd string, limit int64) (map[string]string, error)

	// rscan key_start key_end limit
	// list key-value for (key_start, key_end] reverse
	RScan(keyStart, keyEnd string, limit int64) (map[string]string, error)

	Close() error
}
