FROM node:26-alpine AS frontend

WORKDIR /frontend

COPY frontend/package*.json ./
RUN npm ci

COPY frontend/ .

RUN npm run build

FROM golang:1.26-alpine AS backend

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download -x

COPY . .

COPY --from=frontend /frontend/dist ./frontend/dist

RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w" \
    -o /app/webfeesh \
    ./cmd/web

FROM alpine:latest AS final

RUN addgroup -g 1000 -S webfeesh && \
    adduser -u 1000 -S webfeesh -G webfeesh

WORKDIR /app

RUN mkdir -p /app/data

COPY --from=backend /app/webfeesh /app/webfeesh

RUN chown -R webfeesh:webfeesh /app

USER webfeesh

EXPOSE 4000
ENTRYPOINT ["/app/webfeesh"]
