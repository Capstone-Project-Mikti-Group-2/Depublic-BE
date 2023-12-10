# Golang base image
FROM golang:1.21.5-alpine3.18

# select time zone
RUN ln -sf /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
    echo "Asia/Jakarta" > /etc/timezone

# set working directory
WORKDIR /app

# copy go mod and sum files
COPY go.mod .
COPY go.sum .

# Download and install dependencies
RUN export GOPROXY=https://proxy.golang.org && \
    go mod tidy

# copy the source code
COPY . .

# Build Golang Application
RUN go build -o main cmd/server/main.go

# Remove unnacessary files
RUN rm -rf go.mod go.sum

# Expose port
EXPOSE 8080

# Command to run the executable
CMD ["./main"]