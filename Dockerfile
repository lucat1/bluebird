FROM golang:alpine AS go-builder
WORKDIR /workdir
COPY . .
RUN apk update && apk add --no-cache gcc musl-dev
RUN go get -d -v
RUN go build -ldflags="-extldflags=-static" -o bluebird

FROM node:alpine AS node-builder
WORKDIR /workdir
COPY . .
RUN yarn
RUN yarn build

FROM scratch
COPY --from=go-builder /workdir/bluebird /bluebird
COPY --from=node-builder /workdir/dist /dist
ENTRYPOINT ["/bluebird"]
