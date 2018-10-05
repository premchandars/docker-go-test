FROM golang:latest

RUN mkdir /app
WORKDIR /app
ADD . /app/ 



RUN go build -o ocr src/hello-world.go 

EXPOSE 3001
CMD ["/app/ocr"]
