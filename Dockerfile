# Start from golang base image
FROM golang:1.18 AS builder
# FROM golang:buster as builder

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Generate the prisma data layer
RUN go install github.com/prisma/prisma-client-go
RUN go run github.com/prisma/prisma-client-go generate

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd
RUN echo The old build end here 
# Start a new stage from scratch
# FROM alpine:latest
# FROM golang:buster
FROM alpine AS production
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add curl

# WORKDIR /root/

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder /usr/local/go/ /usr/local/go/
COPY --from=builder /app .
# COPY --from=builder /app/prisma ./seed/prisma
# COPY --from=builder /app/go.mod ./seed/
# COPY --from=builder /app/go.sum ./seed/
# COPY --from=builder /app/main ./seed/
# COPY --from=builder prisma.yaml .

# RUN apk update && apk add go gcc bash musl-dev openssl-dev ca-certificates && update-ca-certificates

# ENV PATH="/usr/local/go/bin:${PATH}"

EXPOSE $PORT


#Command to run the executable
CMD [ "./main" ]
