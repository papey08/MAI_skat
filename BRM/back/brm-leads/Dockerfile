FROM golang:1.21.6

ENV GOPATH=/

WORKDIR /go/src/brm-leads
COPY . .

RUN go mod download
RUN go build -o brm-leads-app cmd/server/main.go

CMD ["./brm-leads-app --docker"]
