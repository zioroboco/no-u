build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/handler handler/*

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: deploy
deploy: clean build
	sls deploy --verbose
