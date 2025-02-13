# Stage 1: Builder
FROM golang AS builder
RUN apt update
RUN apt install -y sudo systemctl
WORKDIR /app
COPY go.mod ./
RUN go mod download
RUN go install github.com/go-task/task/v3/cmd/task@latest
COPY . .
WORKDIR /app/example
RUN task build

# Stage 2: Installer
FROM builder AS installer
RUN task install
RUN task start
RUN task stop
RUN task status
RUN task uninstall