.PHONY: build up down logs shell mysql

build:
	docker-compose build --no-cache

up:
	docker-compose up -d

down:
	docker-compose down

stop:
	docker-compose stop

logs:
	docker-compose logs -f

shell:
	docker-compose exec app /bin/bash

db:
	docker-compose exec db /bin/bash
