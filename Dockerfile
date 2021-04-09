FROM golang:1.16-alpine as builder
LABEL maintainer="Douglas Dennys <douglasdennys45@gmail.com>"
WORKDIR /app
COPY . .
RUN go mod download
RUN go build ./src/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]