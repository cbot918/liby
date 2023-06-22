# connect to db
DOCKER_NETWORK="bridge"
psql:
	docker run -it --rm --network=$(DOCKER_NETWORK) postgres psql -h localhost -U postgres

.PHONY: psql