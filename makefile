build: 
		go build -o bin/currency-microservice

run: build
		./bin/currency-microservice
		
test:
		@go test -v ./...