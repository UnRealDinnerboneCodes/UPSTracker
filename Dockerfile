FROM golang:1.18-alpine as build
LABEL stage=Build
WORKDIR /go/src/app
COPY . .
RUN go get .
RUN go build -o /upstracker

FROM alpine:3
COPY --from=build /upstracker /upstracker
EXPOSE 8090

ENTRYPOINT  ["/upstracker"]