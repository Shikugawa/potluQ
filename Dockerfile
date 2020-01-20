FROM golang:latest as builder
ENV GOPATH=/go
ENV GO111MODULE=on
WORKDIR ${GOPATH}/src/github.com/Shikugawa/potluq
COPY . .
RUN make build

FROM alpine:latest
RUN apk add --update --no-cache ca-certificates tzdata && update-ca-certificates
WORKDIR /app
COPY --from=builder /go/src/github.com/Shikugawa/potluq/dist/potluq .
RUN chmod +x ./potluq
EXPOSE 8000