package repository

import "context"

type KeyValueStore interface {
	Connect() error
	Close() error
	GetValue(ctx context.Context, key string) (*string, error)
	SetValue(ctx context.Context, key, value string) error
	Count(ctx context.Context) (int, error)
}
