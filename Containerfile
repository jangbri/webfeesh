FROM golang:1.26-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download -x

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build \
    -ldflags="-s -w" \
    -o /app/webfeesh \
    ./cmd/web

FROM alpine:latest AS final

RUN apk --update add \
    ca-certificates \
    tzdata \
    && \
    update-ca-certificates

RUN addgroup -g 1000 -S webfeesh && \
    adduser -u 1000 -S webfeesh -G webfeesh

WORKDIR /app

RUN mkdir -p /app/data

COPY --from=builder /app/webfeesh /app/webfeesh

RUN chown -R webfeesh:webfeesh /app

USER webfeesh

EXPOSE 4000
ENTRYPOINT ["/app/webfeesh"]
