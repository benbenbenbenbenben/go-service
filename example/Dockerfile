# Stage 1: Builder
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/go-task/task/v3/cmd/task@latest
COPY . .
WORKDIR /app/example
RUN task build

# Stage 2: Installer
FROM alpine:latest AS installer
WORKDIR /app
COPY --from=builder /app/example/build/example /app/example
RUN apk add --no-cache sudo
RUN chmod +x example
RUN sudo ./example -install
RUN sudo ./example -start
RUN sudo ./example -stop
RUN sudo ./example -status
RUN sudo ./example -uninstall

# Stage 3: Runner (Immediate Mode)
FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/build/example /app/example
RUN chmod +x example
CMD ["./example", "-run"]
