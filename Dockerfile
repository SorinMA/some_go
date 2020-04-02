FROM golang:alpine
WORKDIR ./
COPY . .
EXPOSE 8000
RUN go build  main.go helpers.go server.go 
ENTRYPOINT ./main
