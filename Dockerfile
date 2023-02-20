FROM golang

WORKDIR  /falconEmailServer

COPY . .
RUN go build ./app/falconEmailServer.go
ENTRYPOINT ./falconEmailServer

EXPOSE 3000