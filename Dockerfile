FROM golang

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build
COPY app app
COPY go.mod .
COPY go.sum .
RUN go build -o main app/main.go

# Run app
WORKDIR /app
RUN cp /build/main app
CMD ["/app/app"]
