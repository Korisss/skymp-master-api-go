#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./
RUN go build -o /go/bin/app -v ./cmd/

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
COPY --from=builder /go/src/app/configs/ /configs/
ENTRYPOINT /app
LABEL Name=skympmasterapigo Version=0.0.1
EXPOSE 3000
