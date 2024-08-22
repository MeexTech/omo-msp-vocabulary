.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/vocabulary/concept.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/vocabulary/attribute.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/vocabulary/common.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/vocabulary/entity.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/vocabulary/event.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/vocabulary/graph.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/vocabulary/relation.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/vocabulary/box.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/vocabulary/edge.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/vocabulary/examine.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/vocabulary/import.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/vocabulary/template.proto
