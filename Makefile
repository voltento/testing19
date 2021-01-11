GROUP := voltento
PROJECT := testing19
TAG := ${GROUP}-${PROJECT}

run: build-image
	docker run -it ${TAG}

build-image:
	docker build --build-arg project_dir=${GROUP}/${PROJECT} . -t ${TAG}

test:
	go test ./...