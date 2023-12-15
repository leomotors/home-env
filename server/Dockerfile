FROM golang:1.21-alpine3.19 as builder

# Set destination for COPY
WORKDIR /app

# Copy all source files and download Go modules
COPY . ./
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o ./home-env-server ./cmd/main.go

FROM alpine:3.19 AS runner

LABEL name "home-env-server"

WORKDIR /app
USER guest

# Copy built executable
COPY --chown=guest:guest --from=builder /app/home-env-server ./
# Assets
COPY --chown=guest:guest public ./public

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8939

# Run
CMD ["./home-env-server"]
