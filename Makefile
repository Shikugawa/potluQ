build:
	GOOS=linux CGO_ENABLED=0 go build -o ./dist/potluq -i ./main.go	

build-image:
	docker build . -t shikugawa/potluq