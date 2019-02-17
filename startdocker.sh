docker container stop cc-payment
docker container rm cc-payment
docker run -d --name cc-payment \
    -v $PWD/cfg:/home/paymentserv/cfg \
    -v $PWD/logs:/home/paymentserv/logs \
    cc-payment