FROM golang:1.16.4-alpine3.13 as builder

WORKDIR /src

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY gen/ gen/
COPY cmd/ cmd/
COPY static/ static/
COPY facts.go .
COPY goat_facts.go .
COPY static.go .

RUN go generate

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" ./cmd/goatopsfarm

FROM scratch

WORKDIR /app

COPY --from=builder /src/goatopsfarm /app/goatopsfarm

CMD ["/app/goatopsfarm"]
