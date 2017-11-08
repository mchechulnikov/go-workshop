FROM golang:1.8.3-alpine

EXPOSE 80

WORKDIR /go/src/auth
COPY . .
RUN chmod +x ./start.sh

RUN go install

ENTRYPOINT ./start.sh
