FROM golang


WORKDIR /go/src/app
COPY . .

RUN go mod download -x
RUN go install github.com/githubnemo/CompileDaemon


ENTRYPOINT ["/go/bin/CompileDaemon", "--build", "go build -o app ./server/cmd/crypto-checkout-server", "--command", "./app --port 8082 --host 0.0.0.0"]