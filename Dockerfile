FROM node:latest

# Install bun
RUN curl -fsSL https://bun.sh/install | bash
ENV PATH="/root/.bun/bin:${PATH}"

# Install Go
RUN apt-get update && apt-get install -y golang-go

# Install git
RUN apt-get install -y git

# Set Go environment variables
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH

# Set working directory
WORKDIR /app

# Copy package.json and install dependencies
COPY frontend/package.json ./frontend/
RUN cd frontend && bun install

# Copy Go files
COPY backend ./backend

# Expose ports
EXPOSE 3000 8080 3306

RUN cd backend && go install github.com/air-verse/air@latest

# Start air
CMD ["air", "-c", "/backend/.air.toml"]