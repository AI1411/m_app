exec-schema: ## sqlファイルをPostgresコンテナに流す
	cat ./migrations/*.up.sql > ./migrations/schema.sql
	docker cp migrations/schema.sql db:/schema.sql
	docker exec -it db psql -U myuser -d m_app -f /schema.sql
	docker cp migrations/schema.sql db-test:/schema.sql
	docker exec -it db-test psql -U testuser -d m_app_test -f /schema.sql
	rm ./migrations/schema.sql
# Generate Model
# https://gorm.io/gen/gen_tool.html
# generating all of associations is required multiple execute angenerate_associations/main.go  .
.PHONY: generate-models
generate-models:
	@rm -rf ./internal/domain/query/*.gen.go ./internal/domain/model/*.gen.go
	@go run ./cmd/gormgen/generate_all/main.go
	@go run ./cmd/gormgen/generate_associations/main.go
seed:
	@echo "seeding..."
	@docker cp ./migrations/seed.sql db:/seed.sql
	@docker exec -it db psql -U myuser -d m_app -f /seed.sql
	@echo "seeding done"

.PHONY: install-buf-plugins
install-buf-plugins:
	@echo "Installing buf plugins..."
	@go install github.com/bufbuild/buf/cmd/buf@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@echo "Buf plugins installed successfully"

.PHONY: generate-proto
generate-proto: install-buf-plugins
	@echo "Generating proto files..."
	@buf generate
	@echo "Proto files generated successfully"
