GROUP := voltento
PROJECT := testing19

run: build-image
	docker run -it voltento-testing19

build-image:
	docker build --build-arg project_dir=${GROUP}/${PROJECT} . -t voltento-testing19

test:
	go test ./...