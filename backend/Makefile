exec-schema: ## sqlファイルをPostgresコンテナに流す
	cat ./migrations/*.up.sql > ./migrations/schema.sql
	docker cp migrations/schema.sql db:/schema.sql
	docker exec -it db psql -U myuser -d m_app -f /schema.sql
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