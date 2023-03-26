run:
	go run cmd/main.go
build:
	go build cmd/main.go
docker-build:
	docker build . -t termboard-image
docker-run:
	docker run -it --rm --name termboard termboard-image