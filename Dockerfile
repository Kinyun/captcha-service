FROM golang:1.21.11-alpine as build-env

WORKDIR /app

# Copy the Go modules files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

COPY . .

RUN go mod tidy && go mod download

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM alpine:latest

WORKDIR /app
# Create a non-root user and group
RUN addgroup -S golang && adduser -S gouser -G golang

COPY --from=build-env /go/bin/app /app/app
COPY assets/ /app/assets

# Change ownership to the non-root user
RUN chown gouser:golang app

# Use the non-root user to run the application
USER gouser

CMD ["/app/app"]
