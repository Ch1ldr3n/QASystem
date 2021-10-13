FROM node:14 as frontend
ADD frontend /src
WORKDIR /src
RUN npm install -g cnpm --registry=https://registry.npmmirror.com
RUN cnpm install
RUN cnpm run build

FROM golang:1.17-alpine as backend
ADD backend /go/src/backend
WORKDIR /go/src/backend
RUN GOPROXY=https://goproxy.cn go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -d pkg -o pkg/docs
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk update && apk add build-base
RUN GOPROXY=https://goproxy.cn go build -o /go/bin/qanda ./cmd/qanda

FROM alpine:edge
COPY --from=frontend /src/dist /usr/share/qanda
COPY --from=backend /go/bin/qanda /usr/bin/qanda
ENTRYPOINT [ "/usr/bin/qanda", "-listen", ":80", "-storage", "postgres", "-database", "postgres://postgres:7a5d3540fabb91e4be5b8d60b5c10767@database.bauhinia.secoder.local/postgres?sslmode=disable" ]
EXPOSE 80