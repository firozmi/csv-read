include .env

start:
	mkdir ${LOG_FILE} && touch ${LOG_FILE}/error.log
	docker build -t go-csv-read .
	docker run -d -p ${PORT}:8111 --name go-csv-read -v ${LOG_FILE}:/go/src/app/log/ go-csv-read
	@echo 'Visit http://localhost:${PORT}'

clean:
	-docker stop go-csv-read
	-docker rm go-csv-read
	-docker rmi go-csv-read
	-rm -r ${LOG_FILE}

test:
	go test bitbucket.org/firozmi/csv-read/src/handler

## All targets should have a ## Help text above the target and they'll be automatically collected
## Show help, using auto generator from https://gist.github.com/prwhite/8168133
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)