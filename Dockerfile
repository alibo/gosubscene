
#build stage
FROM golang:1.14-alpine AS builder
WORKDIR /go/src/app
RUN apk add --no-cache git

COPY go.sum go.mod /go/src/app/
RUN go mod download

COPY . /go/src/app
RUN go build -ldflags="-w -s" -o gosubscene

#final stage
FROM alpine:3.11
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /go/src/app/gosubscene /app/gosubscene
ENTRYPOINT /app/gosubscene
LABEL Name=gosubscene Version=0.0.1
EXPOSE 3000
