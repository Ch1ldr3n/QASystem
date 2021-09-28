FROM golang:1.17 as backend
ADD backend /go/src/backend
WORKDIR /go/src/backend
RUN GOPROXY=https://goproxy.cn CGO_ENABLED=0 go build -o /go/bin/qanda ./cmd/qanda

FROM alpine:edge
COPY --from=backend /go/bin/qanda /usr/bin/qanda
ENTRYPOINT [ "/usr/bin/qanda", "-listen", ":80" ]
EXPOSE 80