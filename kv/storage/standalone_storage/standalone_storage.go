package standalone_storage

import (
	"tinykv/kv/config"
	"tinykv/kv/storage"
	"tinykv/kv/storage/engine"
	"tinykv/proto/kvrpcpb"
)

// StandaloneStorage stores data locally
type StandaloneStorage struct {
	engine *engine.Engines
}

func NewStandaloneStorage(conf *config.Config) *StandaloneStorage {
	return &StandaloneStorage{
		engine: engine.NewEngines(engine.CreateDB("", conf), nil, conf.DBPath, ""),
	}
}

func (ss *StandaloneStorage) Start() error {
	return nil
}

func (ss *StandaloneStorage) Stop() error {
	return ss.engine.Close()
}

func (ss *StandaloneStorage) Write(ctx *kvrpcpb.Context, batch []storage.Modify) error {
	var writeBatch engine.WriteBatch
	for _, m := range batch {
		writeBatch.SetCF(engine.CfDefault, m.Key(), m.Value())
	}
	return ss.engine.WriteKV(&writeBatch)
}

func (ss *StandaloneStorage) Reader(ctx *kvrpcpb.Context) (storage.StorageReader, error) {
	txn := ss.engine.Kv.NewTransaction(false)
	return NewStandaloneReader(txn), nil
}
