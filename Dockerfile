FROM golang:alpine 

RUN apk update && apk add git && apk add make
WORKDIR /app
COPY . .

RUN go mod tidy
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go build -o binary

ENTRYPOINT [ "/app/binary" ]