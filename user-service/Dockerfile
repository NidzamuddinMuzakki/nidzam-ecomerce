FROM golang:1.16-alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go build -o user-service

############################
# STEP 2 build a small image
############################
FROM builder
WORKDIR /app
COPY --from=builder /app/user-service /app
CMD ["./user-service"]