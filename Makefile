create:
	docker-compose up -d --build
up:
	docker-compose up -d
down:
	docker-compose down
destroy:
	docker-compose down --rmi all --volumes
bash:
	docker-compose exec app bash