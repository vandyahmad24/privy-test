
create-migration:
	migrate create -ext sql -dir migration $(NAME)	
run-migration:
	migrate -database "mysql://root:root@tcp(localhost:3305)/cake" -path migration up
run:
	go run main.go