FROM golang:1.17 AS builder
  

ENV GO111MODULE=off \
    CGO_ENABLED=0 \
    GOOS=linux \
    GORACH=amd64

WORKDIR /builder
COPY . .
RUN go build -o httpserver .

FROM scratch
COPY --from=builder /builder/httpserver /
EXPOSE 8080
ENTRYPOINT ["/httpserver"] 