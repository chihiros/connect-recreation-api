FROM golang:1.22-bullseye AS docker
WORKDIR /app
COPY /app .
RUN go install github.com/cosmtrek/air@latest
CMD air

FROM golang:1.22-bullseye AS build
WORKDIR /app
COPY /app .
RUN go mod download
RUN GOOS=linux go build -tags timetzdata -mod=readonly -v -o server

FROM ubuntu:latest AS deploy
RUN apt update && apt install -y ca-certificates
COPY --from=build /app/server /server

CMD ["/server"]
