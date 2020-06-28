package standalone_storage

import (
	"tinykv/kv/storage/engine"

	"github.com/Connor1996/badger"
)

type StandaloneReader struct {
	txn *badger.Txn
}

func NewStandaloneReader(txn *badger.Txn) *StandaloneReader {
	return &StandaloneReader{
		txn: txn,
	}
}

func (sr *StandaloneReader) GetCF(cf string, key []byte) ([]byte, error) {
	val, err := engine.GetCFFromTxn(sr.txn, cf, key)
	if err == badger.ErrKeyNotFound {
		return nil, nil
	}
	return val, err
}

func (sr *StandaloneReader) IterCF(cf string) engine.DBIterator {
	return engine.NewCFIterator(cf, sr.txn)
}

func (sr *StandaloneReader) Close() {
	sr.txn.Discard()
}
