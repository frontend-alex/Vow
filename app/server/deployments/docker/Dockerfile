FROM golang:1.26-alpine AS build

WORKDIR /src
COPY go.mod ./
COPY . ./
RUN go build -o /bin/vow-api ./cmd/api

FROM alpine:3.22
COPY --from=build /bin/vow-api /bin/vow-api
EXPOSE 8080
CMD ["/bin/vow-api"]
