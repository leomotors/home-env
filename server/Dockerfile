# syntax=docker/dockerfile:1
FROM golang:1.21-alpine3.18 as builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /home-env-server

FROM alpine:3.18 AS runner

WORKDIR /app

USER guest

# Copy built files
COPY --chown=guest:guest --from=builder /home-env-server ./
# Assets
COPY --chown=guest:guest *.html ./

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD ["./home-env-server"]
