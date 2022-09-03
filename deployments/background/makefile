DOCKER=docker
COMPOSE_TEST_FILE=docker-compose.test.yaml

test-build:
	${DOCKER} compose \
		-f ./${COMPOSE_TEST_FILE} \
		build

test: test-build
	${DOCKER} compose \
		-f ./${COMPOSE_TEST_FILE} \
		up app_test \
		--abort-on-container-exit \
		--exit-code-from app_test

.ONESHELL:
build:
	${DOCKER} compose build
	touch secret.env
	ls -la
	echo "SECRET_TOKEN=${SECRET_TOKEN}" | tee ./secret.env
	echo "MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD}" | tee -a ./secret.env
	ls -la