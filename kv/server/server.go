package server

import (
	"context"
	"tinykv/kv/storage"
	"tinykv/proto/kvrpcpb"
	"tinykv/zaplog"

	"go.uber.org/zap"
)

type Server struct {
	storage storage.Storage
}

func NewServer(storage storage.Storage) *Server {
	return &Server{
		storage: storage,
	}
}

// Raw API.
func (server *Server) RawGet(_ context.Context, req *kvrpcpb.RawGetRequest) (*kvrpcpb.RawGetResponse, error) {
	resp := new(kvrpcpb.RawGetResponse)
	// get reader
	r, err := server.storage.Reader(nil)
	if err != nil {
		resp.Error = err.Error()
		return resp, err
	}
	// get cf
	val, err := r.GetCF(req.GetCf(), req.GetKey())
	if err != nil {
		resp.Error = err.Error()
	}
	if val == nil {
		resp.NotFound = true
	} else {
		resp.Value = val
	}

	return resp, nil
}

func (server *Server) RawScan(_ context.Context, req *kvrpcpb.RawScanRequest) (*kvrpcpb.RawScanResponse, error) {
	resp := new(kvrpcpb.RawScanResponse)

	var res []*kvrpcpb.KvPair
	// get reader
	r, err := server.storage.Reader(nil)
	if err != nil {
		resp.Error = err.Error()
		return resp, err
	}
	// scan cf
	iter := r.IterCF(req.GetCf())
	defer iter.Close()
	iter.Seek(req.GetStartKey())
	for iter.Valid() {
		k := iter.Item().Key()
		v, err := iter.Item().Value()
		if err != nil {
			zaplog.Error("iter item error", zap.Error(err))
			continue
		}
		item := kvrpcpb.KvPair{
			Key:   k,
			Value: v,
		}
		res = append(res, &item)
	}
	if len(res) > 0 {
		resp.Kvs = res
	} else {
		resp.Error = "no result"
	}

	return resp, nil
}

func (server *Server) RawPut(_ context.Context, req *kvrpcpb.RawPutRequest) (*kvrpcpb.RawPutResponse, error) {
	resp := new(kvrpcpb.RawPutResponse)
	var modify []storage.Modify
	modify = append(modify, storage.Modify{
		Data: storage.Put{
			Key:   req.GetKey(),
			Value: req.GetValue(),
			Cf:    req.GetCf(),
		},
	})
	err := server.storage.Write(nil, modify)
	if err != nil {
		resp.Error = err.Error()
	}
	return resp, err
}

func (server *Server) RawDelete(_ context.Context, req *kvrpcpb.RawDeleteRequest) (*kvrpcpb.RawDeleteResponse, error) {
	resp := new(kvrpcpb.RawDeleteResponse)
	var modify []storage.Modify
	modify = append(modify, storage.Modify{
		Data: storage.Delete{
			Key: req.GetKey(),
			Cf:  req.GetCf(),
		},
	})
	err := server.storage.Write(nil, modify)
	if err != nil {
		resp.Error = err.Error()
	}
	return resp, err
}
