FROM golang:1.10 as builder

MAINTAINER zerro "zerrozhao@gmail.com"

WORKDIR $GOPATH/src/github.com/zhs007/cc-payment

COPY ./Gopkg.* $GOPATH/src/github.com/zhs007/cc-payment/

RUN go get -u github.com/golang/dep/cmd/dep \
    && dep ensure -vendor-only -v

COPY . $GOPATH/src/github.com/zhs007/cc-payment

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o paymentserv . \
    && mkdir /home/paymentserv \
    && mkdir /home/paymentserv/cfg \
    && mkdir /home/paymentserv/dat \
    && mkdir /home/paymentserv/logs \
    && cp ./paymentserv /home/paymentserv/

FROM alpine
RUN apk upgrade && apk add --no-cache ca-certificates
WORKDIR /home/paymentserv
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=builder /home/paymentserv /home/paymentserv
CMD ["./paymentserv"]