FROM golang:1.16-alpine as builder
LABEL maintainer="Douglas Dennys <douglasdennys45@gmail.com>"
WORKDIR /app
COPY . .
RUN go mod download && go build ./src/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/.env.example .
RUN cp .env.example .env && rm -r .env.example
EXPOSE 8080
CMD ["./main"]