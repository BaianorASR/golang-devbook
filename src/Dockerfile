    FROM golang:alpine
    RUN mkdir /app
    ADD . /app/
    
    WORKDIR /app
    
    RUN apk update && apk upgrade && apk add --no-cache bash git openssh
    
    RUN go get github.com/jinzhu/gorm
    RUN go get github.com/jinzhu/gorm/dialects/mysql
    RUN go get github.com/gin-gonic/gin
    RUN go get github.com/gin-contrib/cors
    
    
    RUN go build -o main .
    RUN adduser -S -D -H -h /app appuser
    USER appuser
    CMD ["./main"]