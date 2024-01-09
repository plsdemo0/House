test:
	@echo 'Мы сделали Makefile'

up:
	sudo -S docker-compose up --build awesome_project

stop:
	sudo -S docker-compose stop