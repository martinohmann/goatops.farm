FROM golang:1.16.4-alpine3.13 as builder

WORKDIR /src

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY main.go .
COPY tools/ tools/

RUN go generate

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o goatops.farm

FROM scratch

WORKDIR /app

COPY --from=builder /src/goatops.farm /app/goatops.farm
COPY --from=builder /src/goatfacts.json /app/goatfacts.json

CMD ["/app/goatops.farm"] 
