.PHONY: build up down logs shell db

build:
	docker-compose build --no-cache

up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f

shell:
	docker-compose exec app /bin/bash

db:
	docker-compose exec db mysql -uroot -prootpassword myapp