FROM golang:latest

ENV PORT 8000
WORKDIR /app

# Build
COPY clock.go .
RUN go build -o clock .

# Run
CMD ["/app/clock -port $PORT"]