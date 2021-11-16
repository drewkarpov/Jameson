FROM golang:alpine AS builder

WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

RUN go get -d -v
RUN go build -o Jameson ./main.go

ENTRYPOINT ["./Jameson"]
