package repository

type KeyValueStore interface {
	Connect() error
	Close() error
	GetValue(key string) (*string, error)
	SetValue(key, value string) error
}
