api-migrate:
	docker compose exec api sh -c "migrate -database $DB_URL -path $MIGRATION_PATH up"