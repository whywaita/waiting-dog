FROM golang:latest

# build directories
RUN mkdir /app
RUN mkdir /go/src/app
ADD . /go/src/app
WORKDIR /go/src/app

# Build my app
RUN go mod download
RUN go build -o /app/waiting-dog .
CMD ["/app/waiting-dog"]
