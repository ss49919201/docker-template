# syntax=docker/dockerfile:1

FROM golang:1.16-alpine as build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY main.go ./

RUN go build -o /hello-world

FROM golang:1.16-alpine as runtime

WORKDIR /app

COPY --from=build /hello-world /bin/hello-world

CMD [ "/bin/hello-world" ]