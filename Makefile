.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/vocabulary/concept.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/vocabulary/entity.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/vocabulary/graph.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/vocabulary/relation.proto
