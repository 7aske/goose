main=cmd/goose.go
out=build

default_recipe: build

build: $(main)
	go build -o $(out) $(main)

run:
	go run $(main)

.PHONY: install
install: build
	sudo cp $(out)/goose /usr/local/bin/


dep:
	echo Installing dependecies

	go get gopkg.in/yaml.v2
	go get github.com/gorilla/mux
	go get github.com/teris-io/shortid

