postgrespull:
	sudo docker pull postgres

postgres:
	sudo docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=kursUser -e POSTGRES_PASSWORD=kursPswd -d postgres:12-alpine

createbd: 
	sudo docker exec -it postgres12 createdb --username=kursUser --owner=kursUser test
dropdb:
	sudo docker exec -it postgres12 dropdb --username=kursUser test

migrateup: 
	migrate -path migrations/ -database "postgresql://kursUser:kursPswd@localhost:5432/test?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations/ -database "postgresql://kursUser:kursPswd@localhost:5432/test?sslmode=disable" -verbose down

.PHONY: postgres createbd dropdb migrateup migratedown

