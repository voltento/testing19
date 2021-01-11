FROM golang:alpine as builder

ARG project_dir

RUN mkdir $GOPATH/src/project_dir

COPY . $GOPATH/src/project_dir
WORKDIR $GOPATH/src/project_dir

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go build -v -o main app/main.go
RUN cp main /main

FROM golang:alpine as production
# Run app
WORKDIR /app
COPY --from=builder /main /app/app
ENTRYPOINT ["/app/app"]
