protoc --proto_path=proto/metapb/ --go_out=paths=source_relative,plugins=grpc:proto/metapb/ proto/metapb/metapb.proto
protoc --proto_path=proto/eraftpb/ --go_out=paths=source_relative,plugins=grpc:proto/eraftpb/ proto/eraftpb/eraftpb.proto
protoc --proto_path=proto/ --go_out=paths=source_relative,plugins=grpc:proto/ proto/errorpb/errorpb.proto
protoc --proto_path=proto/ --go_out=paths=source_relative,plugins=grpc:proto/ proto/coprocessor/coprocessor.proto
protoc --proto_path=proto/ --go_out=paths=source_relative,plugins=grpc:proto/ proto/kvrpcpb/kvrpcpb.proto
protoc --proto_path=proto/ --go_out=paths=source_relative,plugins=grpc:proto/ proto/raft_serverpb/raft_serverpb.proto
protoc --proto_path=proto/ --go_out=paths=source_relative,plugins=grpc:proto/ proto/tinykvpb/tinykvpb.proto