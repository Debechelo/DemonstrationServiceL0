##
## Build
##
FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o /server .

##
## Deploy
##
FROM gcr.io/distroless/static AS deploy

WORKDIR /

COPY --from=build /server /server

EXPOSE 8080

USER nonroot:nonroot

#CMD ["./main"]
ENTRYPOINT ["./server"]

#docker build . -f docker/Dockerfile --tag demo_servis
#docker run -p 8080:8080 demo_servis