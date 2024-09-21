FROM golang:1.22.7-alpine3.19 AS build

ARG service
ARG main_file

WORKDIR /go/src/app

COPY . .

RUN ls -R /go/src/app/

RUN go env -w CGO_ENABLED=0 && \
    go env -w GO111MODULE="on" && \
    go env -w GOOS=linux && \
    go env -w GOARCH=amd64

RUN go mod download

RUN go build -ldflags "-s -w" -o service $main_file

FROM alpine:3.19 AS final

ARG service_port=3000

EXPOSE $service_port

COPY --from=build /go/src/app/service /service

CMD ["/service"]