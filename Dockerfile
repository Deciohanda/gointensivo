FROM golang:1.20

# WORKDIR /go/src
# ENV PATH="go/bin:${PATH}"
WORKDIR /go/goapp

RUN apt-get update && apt-get install -y librdkafka-dev -y

CMD ["tail", "-f", "/dev/null"]
