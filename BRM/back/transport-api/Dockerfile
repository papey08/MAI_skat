FROM golang:1.21.6

ENV GOPATH=/

WORKDIR /go/src/transport-api
COPY . .

RUN go mod download
RUN go build -o transport-api-app cmd/server/main.go

CMD ["./transport-api-app --docker"]
