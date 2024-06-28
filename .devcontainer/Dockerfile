# .devcontainer/Dockerfile
FROM golang:1.20

# Install dependencies
RUN apt-get update && apt-get install -y \
    git \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /workspace

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod download

# Copy the source code
COPY . .

# Install fresh (optional for live reloading)
RUN go install github.com/gravityblast/fresh@latest

# Expose the application ports
EXPOSE 9001