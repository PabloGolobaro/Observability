FROM golang:1.18 as build

COPY . /src

WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux go build -o order_server ./cmd/api

FROM gcr.io/distroless/static-debian11 AS production

COPY --from=build /src/order_server .

EXPOSE 8080

CMD ["/order_server"]