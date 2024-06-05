FROM golang:1.21.6

ENV GOPATH=/

WORKDIR /go/src/notifications
COPY . .

RUN go mod download
RUN go build -o notifications-app cmd/server/main.go

CMD ["./notifications-app --docker"]
