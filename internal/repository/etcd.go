package repository

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var _ KeyValueStore = (*ETCDRepository)(nil)

type ETCDRepository struct {
	client *clientv3.Client
}

func (e *ETCDRepository) Connect() error {
	var err error
	e.client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	return err
}

func (e *ETCDRepository) Close() error {
	return e.client.Close()
}

func (e *ETCDRepository) GetValue(key string) (*string, error) {
	//TODO implement me
	panic("implement me")
}

func (e *ETCDRepository) SetValue(key, value string) error {
	//TODO implement me
	panic("implement me")
}
