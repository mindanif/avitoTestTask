ifeq ($(POSTGRES_SETUP_TEST),)
	POSTGRES_SETUP_TEST := user=root password=secret dbname=trackerDB host=localhost port=5432 sslmode=disable
endif

.PHONY:  build
build:
	 go build -v ./cmd/app

.PHONY: compose-up
compose-up:
	docker-compose build
	docker-compose up -d postgres

 .PHONY: compose-rm
compose-rm:
	docker-compose down

.DEFAULT_GOAL := build

