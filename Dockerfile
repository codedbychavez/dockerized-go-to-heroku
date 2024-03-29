FROM golang:1.17.6

# Add a work directory
WORKDIR /app

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy app files
COPY . .

# Expose port
EXPOSE 4000

# Start app
CMD go run main.go