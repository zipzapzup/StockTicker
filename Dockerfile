# Build stage
FROM golang:1.20 as buildstage

# Create repository
WORKDIR /app

# Download module
COPY go.mod go.sum ./
RUN go mod download

# Copy Source Code
COPY . ./

# Test
RUN go test cmd/stockticker/*

# Build
RUN CGO_ENABLED=0 GOOS=linux go build cmd/stockticker/stockticker.go

FROM alpine:3.18.2 as lean-image

WORKDIR /app

COPY --from=buildstage /app/stockticker /app/stockticker

# Expose

EXPOSE 8080

# Entrypoint
CMD ["/app/stockticker"]
