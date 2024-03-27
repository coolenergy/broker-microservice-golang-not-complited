FROM alpine:3.18

RUN mkdir /app

COPY brokerApp /app

CMD ["/app/brokerApp"]
