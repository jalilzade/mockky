FROM node:12.18.3-alpine3.12 AS JS_BUILD
WORKDIR /webapp
COPY ./cmd/frontend .
RUN npm install && npm run build

# FROM golang:1.15.1-alpine3.12 AS GO_BUILD
# RUN apk add git
# RUN apk update && apk add build-base
# WORKDIR /go/src/app
# COPY ./cmd/server .
# RUN go get -d -v ./...
# RUN go install -v ./...
# RUN go build -o /go/src/app/bin/server

FROM golang:1.15.1-alpine3.12 AS GO_BUILD
RUN apk update && apk add build-base
COPY cmd/server /server
COPY go.mod go.sum /server/
WORKDIR /server
RUN go build -o /go/bin/server

FROM alpine:3.12.0
COPY --from=JS_BUILD /webapp/build* ./webapp/
COPY --from=GO_BUILD /go/bin/server ./
EXPOSE 8080
CMD ./server
