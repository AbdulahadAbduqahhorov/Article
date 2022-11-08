run:
	go run main.go
migrateup:
	migrate -path ./migrations -database 'postgres://abdulahad:10082018@localhost:5432/uacademy?sslmode=disable' up

migratedown:
	migrate -path ./migrations -database 'postgres://abdulahad:10082018@localhost:5432/uacademy?sslmode=disable' down