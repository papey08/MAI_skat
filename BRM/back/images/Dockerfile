FROM golang:1.21.6

ENV GOPATH=/

WORKDIR /go/src/images
COPY . .

RUN go mod download
RUN go build -o images-app cmd/server/main.go

CMD ["./images-app --docker"]
