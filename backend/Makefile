exec-schema: ## sqlファイルをPostgresコンテナに流す
	cat ./migrations/*.up.sql > ./migrations/schema.sql
	docker cp migrations/schema.sql db:/schema.sql
	docker exec -it db psql -U myuser -d m_app -f /schema.sql
	rm ./migrations/schema.sql