# ============================
# 1) Build stage (Go compile)
# ============================
FROM golang:1.24-bookworm AS builder

# Work inside /app
WORKDIR /app

# Copy Go module files from backend and download dependencies
COPY backend/go.mod backend/go.sum ./backend/
WORKDIR /app/backend
RUN go mod download

# Copy the rest of the backend source code
COPY backend/ ./

# Build the Go binary (Linux amd64)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/server .

# ============================
# 2) Runtime stage
# ============================
FROM debian:bookworm-slim

# Install Chromium for chromedp + basic certs/fonts
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        chromium \
        ca-certificates \
        fonts-liberation \
        curl && \
    rm -rf /var/lib/apt/lists/*

# Create non-root user
RUN useradd -m appuser

# Work inside /app
WORKDIR /app

# Copy compiled Go binary from builder stage
COPY --from=builder /app/server /app/server

# Copy frontend files into the image
COPY frontend/ /app/frontend/

# Env vars used by your Go code
ENV PORT=8080
ENV CHROME_PATH=/usr/bin/chromium

# Expose HTTP port
EXPOSE 8080

# Run as non-root user
USER appuser

# Start the server
CMD ["/app/server"]
