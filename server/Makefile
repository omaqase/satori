# Устанавливаем нужные переменные
PROTOC = buf
PROTOC_INCLUDE = /usr/local/include
GOPATH = $(shell go env GOPATH)

# Сервисы
SERVICES = notification user post comment category

# Директории protobuf
PROTO_DIRS = notification/protobuf user/protobuf post/protobuf comment/protobuf category/protobuf

# Генерация protobuf для всех сервисов
.PHONY: generate

generate: generate-notification generate-user generate-post generate-comment generate-category generate-api generate-swagger

# Генерация protobuf для notification сервиса
generate-notification:
	$(PROTOC) generate --template buf.gen.yaml notification/protobuf/notification.proto

# Генерация protobuf для user сервиса
generate-user:
	$(PROTOC) generate --template buf.gen.yaml user/protobuf/user.proto

# Генерация protobuf для post сервиса
generate-post:
	$(PROTOC) generate --template buf.gen.yaml post/protobuf/post.proto

# Генерация protobuf для comment сервиса
generate-comment:
	$(PROTOC) generate --template buf.gen.yaml comment/protobuf/comment.proto

# Генерация protobuf для category сервиса
generate-category:
	$(PROTOC) generate --template buf.gen.yaml category/protobuf/category.proto

# Генерация gRPC Gateway (API)
generate-api:
	$(PROTOC) generate --template buf.gen.yaml notification/protobuf/notification.proto
	$(PROTOC) generate --template buf.gen.yaml user/protobuf/user.proto
	$(PROTOC) generate --template buf.gen.yaml post/protobuf/post.proto
	$(PROTOC) generate --template buf.gen.yaml comment/protobuf/comment.proto
	$(PROTOC) generate --template buf.gen.yaml category/protobuf/category.proto

# Генерация Swagger UI
generate-swagger:
	$(PROTOC) generate --template buf.gen.yaml --swagger_out=logtostderr=true:web/frontend/src/swagger-ui/ notification/protobuf/notification.proto
	$(PROTOC) generate --template buf.gen.yaml --swagger_out=logtostderr=true:web/frontend/src/swagger-ui/ user/protobuf/user.proto
	$(PROTOC) generate --template buf.gen.yaml --swagger_out=logtostderr=true:web/frontend/src/swagger-ui/ post/protobuf/post.proto
	$(PROTOC) generate --template buf.gen.yaml --swagger_out=logtostderr=true:web/frontend/src/swagger-ui/ comment/protobuf/comment.proto
	$(PROTOC) generate --template buf.gen.yaml --swagger_out=logtostderr=true:web/frontend/src/swagger-ui/ category/protobuf/category.proto
