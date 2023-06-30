# Get golang base image
FROM golang:alpine

# Add Maintainer info
LABEL maintainer="Viet"

# install git to install golang dependencies
RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

# Setup working directory
RUN mkdir /app
WORKDIR /app/src/demoproject/

# Copy the source into working directory inside Docker
COPY . .
COPY .env .

RUN cd ./src/demoproject/ \
    && go get -d -v ./... \
    && go install -v ./...\
    && go get -v github.com/githubnemo/CompileDaemon \
    && go install -v github.com/githubnemo/CompileDaemon

ENV PATH="${PATH}:${GOPATH}/bin"
ENV GO111MODULE=on
# enable hot-reload
# Setup hot-reload for dev stage
CMD CompileDaemon --build="go build -buildvcs=false -a -installsuffix cgo -o main ." --command=./main

# Uncomment below code if dont use hot-reload
# # Build the Go app
# RUN go build -o /build

# # Expose port
# EXPOSE 1324

# # RUN the executable
# CMD ["/build"]
