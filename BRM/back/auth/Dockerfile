FROM golang:1.21.6

ENV GOPATH=/

WORKDIR /go/src/auth
COPY . .

RUN go mod download
RUN go build -o auth-app cmd/server/main.go

CMD ["./auth-app --docker"]
