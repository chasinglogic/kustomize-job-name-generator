FROM go:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o kustomize-job-name-generator .

ENTRYPOINT ["/app/kustomize-job-name-generator"]