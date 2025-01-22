FROM golang:1.23.5-alpine
RUN apk update
RUN apk add --no-cache bash make
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
COPY . ./
RUN go build -o server ./cmd/api/main.go
EXPOSE 8081