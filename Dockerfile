FROM golang:1.24.5-alpine3.22

WORKDIR /app

# Update OS packages to patch vulnerabilities
RUN apk update && apk upgrade

COPY . .

RUN go build -o go-devops-app .

CMD ["./go-devops-app"]

EXPOSE 8080