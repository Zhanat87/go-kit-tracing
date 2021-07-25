FROM golang:alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN apk add git libc-dev gcc vim && go mod tidy && go build -o main .
CMD ["/app/main"]
