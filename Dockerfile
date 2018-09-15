FROM golang:latest as builder
WORKDIR /go/src/github.com/creepypasta-club/creepypasta-backend
COPY . .
RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/backend .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/creepypasta-club/creepypasta-backend/bin/backend .
EXPOSE 9000
ENTRYPOINT ["./backend"]
