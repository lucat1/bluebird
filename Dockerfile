FROM golang:alpine AS go-builder
WORKDIR /workdir
COPY . .
RUN apk add --no-cache gcc musl-dev ca-certificates librsvg-dev cairo-dev pkgconfig font-misc-misc
ENV CGO_CFLAGS_ALLOW=".*"
ENV CGO_LDFLAGS_ALLOW=".*"
RUN go get -d -v
RUN go build -o bluebird

FROM node:alpine AS node-builder
WORKDIR /workdir
COPY . .
RUN yarn
RUN yarn build

FROM alpine
RUN apk add --no-cache ca-certificates librsvg cairo
COPY --from=go-builder /workdir/bluebird /bluebird
COPY --from=node-builder /workdir/dist /dist
ENTRYPOINT ["/bluebird"]
