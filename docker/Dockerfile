##
## Build
##
FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./index ./index
COPY ./internal ./internal

# Установка statik
RUN apk add --no-cache git
RUN go get github.com/rakyll/statik

WORKDIR /app/cmd

# Выполнение statik
RUN statik -src=/app

RUN CGO_ENABLED=0 GOOS=linux go build -o /server .

##
## Deploy
##
FROM gcr.io/distroless/static AS deploy

WORKDIR /app

COPY --from=build ./server ./server
COPY --from=build app/index ./index

EXPOSE 8080

USER nonroot:nonroot

#CMD ["./main"]
ENTRYPOINT ["./server"]

#docker build . -f docker/Dockerfile --tag demo_servis
#docker run -p 8080:8080 demo_servis