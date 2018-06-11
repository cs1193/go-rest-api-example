FROM golang:1.10

ARG buildno

RUN echo "Build number: $buildno"

RUN mkdir -p /usr/go/src/app
RUN curl https://glide.sh/get | sh

ENV GOPATH /usr/go

WORKDIR /usr/go/src/app

ADD glide.yaml glide.yaml
ADD glide.lock glide.lock

RUN glide install

COPY . .

RUN go build source/*.go

CMD ["./app"]

EXPOSE 8000
