build:
	@echo 'Building inventory service...'
	go build -ldflags='-s' -o=./bin/inventory

run: build
	./bin/combustion
