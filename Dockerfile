FROM golang:1.21.5 as build
WORKDIR /go/api
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o data-collector-server ./api/

FROM alpine:latest as api
COPY --from=build /go/api .
EXPOSE 3000
CMD ["./data-collector-server"]
