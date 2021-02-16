# Stage 1
FROM golang:1.15.8 as builder

LABEL mainteiner="Matheus Cumpian <matheus.cumpian@hotmail.com>"

RUN apt-get update && apt-get install git
RUN mkdir /build 
ADD . /build/
WORKDIR /build
RUN go get github.com/google/wire/cmd/wire
RUN wire
RUN go mod download
RUN go build -o server .

# Stage 2
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/ /app/
WORKDIR /app
CMD ["./server"]

EXPOSE 8080