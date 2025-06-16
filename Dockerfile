FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o kustomize-job-name-generator .

FROM alpine:latest
COPY --from=builder /app/kustomize-job-name-generator /kustomize-job-name-generator
ENTRYPOINT ["/kustomize-job-name-generator"]