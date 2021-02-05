FROM golang:1.10
ENV GOPATH=/go
WORKDIR /go
RUN mkdir -p src/github.com/egamorim/star-wars-planets
COPY . /go/src/github.com/egamorim/star-wars-planets
WORKDIR /go/src/github.com/egamorim/star-wars-planets
RUN go get -d -v ./...
RUN make build
FROM debian:jessie-slim
RUN apt-get update && apt-get install ca-certificates -y && rm -rf /var/lib/apt/lists/*
WORKDIR /root/
COPY --from=0 /go/src/github.com/egamorim/star-wars-planets .
CMD ["./star-wars"]