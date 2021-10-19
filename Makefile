BINARY_NAME=gwfDemoApp

build:
	@go mod vendor
	@echo "Building GWF..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "GWF built!"

run: build
	@echo "Starting GWF..."
	@./tmp/${BINARY_NAME} &
	@echo "GWF started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping GWF..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped GWF!"

restart: stop start

sync:
	@rm -rf ./vendor
	@go mod vendor
