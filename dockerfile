FROM golang:1.23

WORKDIR /portfolio

COPY . .

RUN go build -o myapp .

EXPOSE 80

CMD [ "./myapp" ]
