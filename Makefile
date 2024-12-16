
createContainer:
	sudo docker run --name postgres16 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:16-alpine

postgresScript:
	cat bankly_table_schema.sql | sudo docker exec  -i postgres16 psql -U root simple_bank

createdb:
	sudo docker exec -it postgres16 createdb --username=root --owner=root simple_bank

dropdb:
	sudo docker exec -it postgres16 dropdb simple_bank

.PHONY: createdb dropdb postgres postgresScript createContainer
