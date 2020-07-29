FROM golang:1.13.8-alpine

RUN mkdir /app
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

EXPOSE 8000
CMD ["go", "run", "main.go"]