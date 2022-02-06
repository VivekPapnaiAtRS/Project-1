FROM golang:1.16-alpine

WORKDIR /server

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /project-1

EXPOSE 8080

CMD [ "/project-1" ]