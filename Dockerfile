FROM golang:1.24.5

WORKDIR /app

# Install air properly
ENV GOBIN=/usr/local/bin
RUN go install github.com/air-verse/air@latest

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Set up air configuration
CMD ["air"]

