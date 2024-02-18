package repository

import (
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

func (e *ETCDRepository) GetValue(key string) (*string, error) {
	_ = key
	panic("implement me")
}

func (e *ETCDRepository) SetValue(key, value string) error {
	_ = key
	_ = value
	panic("implement me")
}
