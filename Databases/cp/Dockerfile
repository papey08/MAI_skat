FROM golang:1.19

ENV GOPATH=/

WORKDIR .
COPY . .

RUN go mod download
RUN go build -o fitness-club-app main.go

CMD ["./fitness-club-app"]
