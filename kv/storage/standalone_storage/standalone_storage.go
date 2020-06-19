package standalone_storage

import (
	"tinykv/kv/config"
	"tinykv/kv/storage"
	"tinykv/proto/kvrpcpb"
)

type StandaloneStorage struct{}

func NewStandaloneStorage(conf *config.Config) *StandaloneStorage {
	return nil
}

func (StandaloneStorage) Start() error {
	panic("implement me")
}

func (StandaloneStorage) Stop() error {
	panic("implement me")
}

func (StandaloneStorage) Write(ctx *kvrpcpb.Context, batch []storage.Modify) error {
	panic("implement me")
}

func (StandaloneStorage) Reader(ctx *kvrpcpb.Context) (storage.StorageReader, error) {
	panic("implement me")
}
