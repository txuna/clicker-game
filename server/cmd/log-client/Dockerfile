FROM golang:alpine AS builder

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY pkg/ /build/pkg
COPY internal /build/internal
COPY cmd/login .
RUN go build -o main

WORKDIR /app
RUN cp /build/main .

FROM alpine
COPY --from=builder /app/main .

ENTRYPOINT [ "/main" ]