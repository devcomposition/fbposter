FRONT_END_BINARY=frontApp

## start: start the front end
start:
	@echo "Starting front end..."
	@cd frontend-service && go build -o ${FRONT_END_BINARY} ./cmd/web
	@cd frontend-service && ./${FRONT_END_BINARY} &

## stop: stop the front end
stop:
	@echo "Stopping front end..."
	@-pkill -f "./${FRONT_END_BINARY}"
	@echo "Stopped front end!"

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@cd frontend-service && go clean
	@cd frontend-service && rm -f ${FRONT_END_BINARY}

## help: displays help
help: Makefile
	@echo "Choose a command: "
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'