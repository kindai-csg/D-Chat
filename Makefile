init:
	cp pre-commit ./.git/hooks/pre-commit
up:
	docker-compose build
	docker-compose up -d
down:
	docker-compose down
log:
	docker-compose logs -f
