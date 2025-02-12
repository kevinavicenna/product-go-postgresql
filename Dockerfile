FROM golang:latest AS buildstage
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o test main.go

# COPY models
# COPY stoa
FROM alpine
RUN adduser -D kevin
USER kevin:kevin

COPY --from=buildstage /src/.env /app/.env
COPY --from=buildstage /src/test /app/test

WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["/app/test"]
