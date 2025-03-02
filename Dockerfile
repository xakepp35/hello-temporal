FROM golang:1.24 AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o . ./...

FROM scratch
WORKDIR /
COPY --from=builder /build/app .
CMD ["/app"]
