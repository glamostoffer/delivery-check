nats_stream_init:
	docker run -p 4222:4222 -p 8222:8222 nats-streaming

postgres_init:
	docker run --name postgres -p 5433:5432 -e POSTGRES_PASSWORD=password -e POSTGRES_USER=root -d postgres:16

create_db:
	docker exec -it postgres createdb --username=root --owner=root delivery-go

migrate_up:
	migrate -path migrations -database "postgresql://root:password@localhost:5433/delivery-go?sslmode=disable" -verbose up

migrate_down:
	migrate -path migrations -database "postgresql://root:password@localhost:5433/delivery-go?sslmode=disable" -verbose down
