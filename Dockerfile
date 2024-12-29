FROM golang:1.23.4-alpine3.20 AS base

FROM base AS build
WORKDIR /src
COPY . .
RUN go mod download
RUN go build -o javacode cmd/app/main.go

FROM base AS app
WORKDIR /app
COPY --from=build /src/javacode .
COPY configs ./configs
CMD ["/app/javacode"]
