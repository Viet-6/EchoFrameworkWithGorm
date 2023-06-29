# Get golang base image
FROM golang:alpine

# Add Maintainer info
LABEL maintainer="Viet"

# install git to install golang dependencies
RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

# Setup working directory
RUN mkdir /app
WORKDIR /app

# Copy the source into working directory inside Docker
COPY . .
COPY .env

# Run all the depedency
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# enable hot-reload
# Setup hot-reload for dev stage
# RUN go get github.com/githubnemo/CompileDaemon
# RUN go get -v golang.org/x/tools/gopls

# ENTRYPOINT CompileDaemon --build="go build -a -installsuffix cgo -o main ." --command=./main

# Comment below code if use hot-reload
# Build the Go app
RUN go build -o /build

# Expose port
EXPOSE 1324

# RUN the executable
CMD ["/build"]
