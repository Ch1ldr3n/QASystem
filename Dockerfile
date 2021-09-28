FROM golang:1.17 as backend
ADD backend /go/src/backend
WORKDIR /go/src/backend
RUN GOPROXY=https://goproxy.cn go build -o /go/bin/qanda ./cmd/qanda

FROM gcr.io/distroless/base-debian10
COPY --from=backend /go/bin/qanda /usr/bin/qanda
ENTRYPOINT [ "/usr/bin/qanda", "-listen", ":8080" ]