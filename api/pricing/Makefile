build:
	@echo 'Building pricing service...'
	go build -ldflags='-s' -o=./bin/pricing ./cmd/api

run: build
	./bin/pricing
