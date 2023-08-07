FROM golang:1.19.2-alpine as build-env

ENV CGO_ENABLE 0

WORKDIR /go/src/github.com/RodolfoBonis/go_boilerplate/
ADD . /go/src/github.com/RodolfoBonis/go_boilerplate/

RUN go build -o go_boilerplate -v /go/src/github.com/RodolfoBonis/go_boilerplate/

COPY . ./

FROM alpine:3.15

WORKDIR /go/src/github.com/RodolfoBonis/go_boilerplate/

COPY --from=build-env /go/src/github.com/RodolfoBonis/go_boilerplate/go_boilerplate /

CMD ["/go_boilerplate"]

EXPOSE 8000