run: build-image
	docker run -it voltento-testing19

build-image:
	docker build . -t voltento-testing19

test:
	go test ./...