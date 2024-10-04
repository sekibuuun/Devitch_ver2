FROM golang:latest

# Set GOPATH
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# Install Air using go install
RUN go install github.com/air-verse/air@latest

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Expose port 8080
EXPOSE 8080

# Start Air for hot-reloading
CMD ["air", "-c", ".air.toml"]
