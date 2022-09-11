FROM golang:1.16-alpine AS builder

WORKDIR /app
COPY ./go.mod .
RUN go mod download

COPY ./ ./

RUN go build -o /{{ .ProjectName }}

FROM gcr.io/distroless/base-debian10 AS runner

WORKDIR /

COPY --from=builder /{{ .ProjectName }} /{{ .ProjectName }}

USER nonroot:nonroot

ENTRYPOINT ["/{{ .ProjectName }}"]