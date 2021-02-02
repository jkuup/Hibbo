FROM golang:latest
RUN mkdir /app
COPY . /app
WORKDIR /app
ENV GOPROXY="https://goproxy.io,direct"
RUN go mod download
RUN go build -o pibbo .
EXPOSE 8080
CMD ["/app/pibbo"]
