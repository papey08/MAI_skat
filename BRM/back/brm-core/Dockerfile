FROM golang:1.21.6

ENV GOPATH=/

WORKDIR /go/src/brm-core
COPY . .

RUN go mod download
RUN go build -o brm-core-app cmd/server/main.go

CMD ["./brm-core-app --docker"]
