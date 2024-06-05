FROM golang:1.21.6

ENV GOPATH=/

WORKDIR /go/src/brm-ads
COPY . .

RUN go mod download
RUN go build -o brm-ads-app cmd/server/main.go

CMD ["./brm-ads-app --docker"]
