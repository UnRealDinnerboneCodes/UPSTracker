FROM golang:1.19-alpine as build
LABEL stage=Build
WORKDIR /go/src/app
COPY ./src .
RUN go get .
RUN go build -o /upstracker

FROM alpine:3
COPY --from=build /upstracker /upstracker
EXPOSE 8080

ENTRYPOINT  ["/upstracker"]