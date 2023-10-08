up:
	docker compose build
	docker compose up

down:
	docker compose down

create-migration $(NAME):
	migrate create -ext sql -dir ./migration/ -seq $(NAME)

migrate-up $(v):
	migrate -path ./migration/ -database "postgresql://dev:secret@localhost:5432/plantdiagnosis?sslmode=disable" -verbose up $(v)

migrate-down $(v):
	migrate -path ./migration/ -database "postgresql://dev:secret@localhost:5432/plantdiagnosis?sslmode=disable" -verbose down $(v)

migrate-force $(v):
	migrate -path ./migration/ -database "postgresql://dev:secret@localhost:5432/plantdiagnosis?sslmode=disable" -verbose force $(v)

migrate-version:
	migrate -path ./migration/ -database "postgresql://dev:secret@localhost:5432/plantdiagnosis?sslmode=disable" -verbose version
