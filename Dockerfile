FROM node:9.5.0 as client-builder
WORKDIR /client
COPY ./client .
RUN npm install && npm run build

FROM golang:1.9.2 as builder
WORKDIR /go/src/github.com/tarkalabs/embedded_resources
RUN go get -u github.com/golang/dep/cmd/dep
COPY . .
RUN dep ensure
COPY --from=client-builder /client/build ./client/build
RUN go get github.com/rakyll/statik
RUN go generate
RUN make linux

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/tarkalabs/embedded_resources/embedded_resources embedded_resources
CMD ["./embedded_resources"]