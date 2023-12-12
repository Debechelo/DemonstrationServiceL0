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

WORKDIR /app/cmd

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