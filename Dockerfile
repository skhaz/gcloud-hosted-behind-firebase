FROM golang:1.21-bullseye
WORKDIR /opt
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o app

FROM gcr.io/distroless/static-debian12
COPY --from=0 /opt/app /
CMD ["/app"]
