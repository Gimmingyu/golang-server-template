## Build
FROM golang:1.19-alpine AS build

WORKDIR /server

COPY . .
RUN apk add --update make gcc g++ git
RUN go mod download

RUN make build

## Deploy
FROM build AS release

WORKDIR /server

EXPOSE 8080

ENTRYPOINT ["./build/appd"]
