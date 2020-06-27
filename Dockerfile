FROM golang:1.13.5
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get github.com/gorilla/mux
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/stretchr/testify/assert
RUN go get go.uber.org/zap
RUN go build -o main .
CMD ["/main"]