package repository

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var _ KeyValueStore = (*ETCDRepository)(nil)

type ETCDRepository struct {
	client *clientv3.Client
}

func (e *ETCDRepository) Connect() error {
	var err error
	e.client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: time.Second,
	})
	return err
}

func (e *ETCDRepository) Close() error {
	return e.client.Close()
}

func (e *ETCDRepository) GetValue(ctx context.Context, key string) (*string, error) {
	res, err := e.client.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if res.Count != 1 {
		return nil, fmt.Errorf("etcd returned %d keys instead of 1", res.Count)
	}
	val := string(res.Kvs[0].Key)
	return &val, nil
}

func (e *ETCDRepository) SetValue(ctx context.Context, key, value string) error {
	_, err := e.client.Put(ctx, key, value)
	return err
}

func (e *ETCDRepository) Count(_ context.Context) (int, error) {
	// TODO implement me
	panic("implement me")
}
