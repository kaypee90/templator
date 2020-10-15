# Compile stage
FROM golang:latest AS build-env

ADD . /templator_dev
WORKDIR /templator_dev

RUN go build -o /templator

# Final stage
FROM debian:buster

EXPOSE 9898

WORKDIR /
COPY --from=build-env /templator /

CMD ["/templator"]