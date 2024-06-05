FROM golang:1.21.6

ENV GOPATH=/

WORKDIR /go/src/stats
COPY . .

RUN go mod download
RUN go build -o stats-app cmd/server/main.go

CMD ["./stats-app --docker"]
