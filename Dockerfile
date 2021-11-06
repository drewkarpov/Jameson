FROM golang:1.17-buster AS build

ENV GOPATH=/
WORKDIR /src/
COPY ./ /src/

# build go app
RUN go mod download; CGO_ENABLED=0 go build -o /Jameson ./main.go


FROM alpine:latest

COPY --from=build /Jameson /Jameson
COPY ./configs/ /configs/

CMD ["/Jameson"]