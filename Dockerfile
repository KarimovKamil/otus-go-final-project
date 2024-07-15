FROM golang:1.22-alpine3.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /app/build/anti-brute-force-app /app/internal/app/app.go

EXPOSE 9012

CMD [ "/app/build/anti-brute-force-app" ]
