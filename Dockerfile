FROM golang:1.16-alpine AS builder

WORKDIR /app
COPY ./go.mod .
RUN go mod download

COPY ./ ./

RUN go build -o /main

FROM gcr.io/distroless/base-debian10 AS runner

WORKDIR /

COPY --from=builder /main /main

USER nonroot:nonroot

ENTRYPOINT ["/main"]