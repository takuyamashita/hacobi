api-migrate-up:
	@ read -p "How many migration you wants to perform (default value: [all]): " MIGATION_COUNT; \
	docker compose exec -e MIGATION_COUNT=$${MIGATION_COUNT} api sh -c '$$(eval echo migrate -database \$${DB_URL} -path \$${MIGRATION_PATH} \$${MIGATION_FILE_NAME} up \$${MIGATION_COUNT})'

api-migrate-down:
	@ read -p "How many migration you wants to perform (default value: [all]): " MIGATION_COUNT; \
	docker compose exec -e MIGATION_COUNT=$${MIGATION_COUNT} api sh -c '$$(eval echo migrate -database \$${DB_URL} -path \$${MIGRATION_PATH} \$${MIGATION_FILE_NAME} down \$${MIGATION_COUNT})'

api-migrate-create:
	@ read -p "Please provide name for the migration: " MIGATION_FILE_NAME; \
	docker compose exec -e MIGATION_FILE_NAME=$${MIGATION_FILE_NAME} api sh -c '$$(eval echo migrate create -ext sql -dir \$${MIGRATION_PATH} \$${MIGATION_FILE_NAME})'
	sudo chown -R $$USER:$$USER ./src/api/migration

api-migrate-force:
	@ read -p "Which version apply dirty=false: " FORCE_VERSION; \
	docker compose exec -e FORCE_VERSION=$${FORCE_VERSION} api sh -c '$$(eval echo migrate -database \$${DB_URL} -path \$${MIGRATION_PATH} \$${MIGATION_FILE_NAME} force \$${FORCE_VERSION})'

api-migrate-refresh: api-migrate-down api-migrate-up

api-test:
	docker compose exec api sh -c 'go test --count 1 ./pkg/...'

front-fmt:
	docker compose exec front sh -c 'npm run fmt'