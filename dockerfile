FROM golang:1.19.2-alpine as build-env

ENV CGO_ENABLE 0

WORKDIR /go/src/github.com/{{cookiecutter.github_username}}/{{cookiecutter.package_name}}/
ADD . /go/src/github.com/{{cookiecutter.github_username}}/{{cookiecutter.package_name}}/

RUN go build -o go_boilerplate -v /go/src/github.com/{{cookiecutter.github_username}}/{{cookiecutter.package_name}}/

COPY . ./

FROM alpine:3.15

WORKDIR /go/src/github.com/{{cookiecutter.github_username}}/{{cookiecutter.package_name}}/

COPY --from=build-env /go/src/github.com/{{cookiecutter.github_username}}/{{cookiecutter.package_name}}/go_boilerplate /

CMD ["/go_boilerplate"]

EXPOSE 8000