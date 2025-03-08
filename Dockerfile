FROM cgr.dev/chainguard/go AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./app cmd/sky-controller/main.go

FROM cgr.dev/chainguard/static

COPY --from=builder /build/app /app

ENTRYPOINT ["/app"]
