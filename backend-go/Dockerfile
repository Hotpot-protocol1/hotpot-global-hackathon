FROM golang:1.20

WORKDIR $GOPATH/src/github.com/Hotpot-protocol1/hotpot-global

COPY go.mod .
COPY go.sum .

RUN go mod tidy
RUN go mod download

COPY . .
RUN go mod tidy

RUN go build -o hotpot-backend -v main.go