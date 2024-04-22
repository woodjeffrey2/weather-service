# use official Golang image
FROM golang:1.22.2-alpine

ENV CGO_ENABLED=0

# set working directory
WORKDIR /app

# Copy the source code
COPY . .

# Download and install the dependencies
RUN go mod download

#EXPOSE the http port
EXPOSE 8080

# Run the api server
CMD ["go", "run", "src/server/main.go"]
