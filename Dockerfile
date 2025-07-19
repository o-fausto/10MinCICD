FROM golang:1.23-alpine

WORKDIR /app

# Update OS packages to patch vulnerabilities
RUN apk update && apk upgrade

COPY . .

RUN go build -o go-devops-app .

CMD ["./go-devops-app"]

EXPOSE 8080