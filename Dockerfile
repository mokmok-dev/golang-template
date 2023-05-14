FROM golang:1.20 as builder

WORKDIR /go/src/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o app ./cmd/app

FROM gcr.io/distroless/static-debian11:nonroot as runner

WORKDIR /

COPY --from=builder /go/src/app/app /app

ENTRYPOINT ["/app"]
