main=cmd/goose.go
out=build

default_recipe: build

.PHONY: build
build: $(main)
	go build  -o $(out)/goose $(main)

run:
	go run $(main)

.PHONY: install
install: build
	sudo cp $(out)/goose /usr/local/bin/

.PHONY:client
client:
	cd ./web && npm install && npm run build

dep:
	go get gopkg.in/yaml.v2
	go get github.com/gorilla/mux
	go get github.com/gorilla/handlers
	go get github.com/teris-io/shortid
	go get github.com/dgrijalva/jwt-go

