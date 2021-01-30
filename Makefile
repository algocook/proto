recipes:
	protoc -I . --go_out=plugins=grpc:. recipes/recipes.proto

.PHONY: recipes