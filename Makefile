build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: deploy
deploy: clean build
	sls deploy --verbose
