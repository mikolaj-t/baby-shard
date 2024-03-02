package repository

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/rs/zerolog/log"
)

var _ KeyValueStore = (*BasicKV)(nil)

type BasicKV struct {
	store      sync.Map
	count      int
	countMutex sync.RWMutex
}

func (b *BasicKV) Connect() error {
	return nil
}

func (b *BasicKV) Close() error {
	return nil
}

func (b *BasicKV) GetValue(_ context.Context, key string) (*string, error) {
	val, ok := b.store.Load(key)
	if !ok {
		return nil, errors.New("not found")
	}

	valString, ok := val.(string)
	if !ok {
		return nil, fmt.Errorf("type assertion for key %q failed", key)
	}
	return &valString, nil
}

func (b *BasicKV) SetValue(_ context.Context, key, value string) error {
	b.countMutex.Lock()
	b.count++
	b.countMutex.Unlock()

	b.store.Store(key, value)
	log.Debug().Msgf("set key %s to value %s", key, value)
	return nil
}

func (b *BasicKV) Count(_ context.Context) (int, error) {
	b.countMutex.RLock()
	count := b.count
	b.countMutex.RUnlock()
	return count, nil
}
