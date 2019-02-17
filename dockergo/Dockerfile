FROM golang:alpine as builder

COPY ./src /src
WORKDIR /src

RUN apk update \
&& apk add git gcc musl-dev\
&& go build \
&& rm -rf /var/cache/apk/*

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /src/dockergo /app
CMD ["./dockergo"]
EXPOSE 9090