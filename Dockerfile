FROM ubuntu
WORKDIR /app
RUN apt update && apt upgrade -y && apt install -y golang
COPY main.go .
RUN go build -o reflector main.go
EXPOSE 8081
ENTRYPOINT [ "./reflector" ]
