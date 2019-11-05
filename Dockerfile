# Stage 1: Builder Container
ARG GO_VERSION=1.13

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk add --update --no-cache ca-certificates git make

RUN mkdir -p /workspace
WORKDIR /workspace
COPY go.* /workspace/

ENV GOFLAGS="-mod=readonly"
RUN go mod download

COPY . /workspace
RUN make build

# Stage 2: Runtime Container
FROM alpine:3.10.1

RUN addgroup -S gourd && adduser -S gourd -G gourd
RUN apk add --update --no-cache ca-certificates

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /workspace/bin/* /usr/local/bin/

RUN mkdir -p /run/gourd
RUN touch /run/gourd/gourdd.sock
RUN chown gourd:gourd /run/gourd /usr/local/bin/gourdd

USER gourd
EXPOSE 8000
CMD ["/usr/local/bin/gourdd"]