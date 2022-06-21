FROM golang:latest AS builder

COPY . /build

WORKDIR /build

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

RUN mv ./app /app

FROM golang:alpine

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app ./

EXPOSE 3000

CMD [ "./app" ]
