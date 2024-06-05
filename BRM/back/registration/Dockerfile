FROM golang:1.21.6

ENV GOPATH=/

WORKDIR /go/src/registration
COPY . .

RUN go mod download
RUN go build -o registration-app cmd/server/main.go

CMD ["./registration-app --docker"]
