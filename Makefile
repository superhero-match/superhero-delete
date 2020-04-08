prepare:
	go mod download

run:
	go build -o bin/main cmd/api/main.go
	./bin/main

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bin/main cmd/api/main.go
	chmod +x bin/main

dkb:
	docker build -t superhero-delete .

dkr:
	docker run -p "3300:3300" -p "8140:8140" superhero-delete

launch: dkb dkr

api-log:
	docker logs api -f

rmc:
	docker rm -f $$(docker ps -a -q)

rmi:
	docker rmi -f $$(docker images -a -q)

clear: rmc rmi

api-ssh:
	docker exec -it api /bin/bash

db-ssh:
	docker exec -it db /bin/bash

PHONY: prepare build dkb dkr launch api-log rmc rmi clear