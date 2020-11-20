up:
	cp pre-commit ./.git/hooks/pre-commit
	docker-compose build
	docker-compose up -d
down:
	docker-compose down
log:
	docker-compose logs -f
test:
	docker-compose exec server go test -v ./test
