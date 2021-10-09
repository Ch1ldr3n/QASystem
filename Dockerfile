FROM node:14 as frontend
ADD frontend /src
WORKDIR /src
RUN npm install
RUN npm run build

FROM golang:1.17 as backend
ADD backend /go/src/backend
WORKDIR /go/src/backend
RUN GOPROXY=https://goproxy.cn go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -d pkg -o pkg/docs
RUN GOPROXY=https://goproxy.cn CGO_ENABLED=0 go build -o /go/bin/qanda ./cmd/qanda

FROM alpine:edge
COPY --from=frontend /src/dist /usr/share/qanda
COPY --from=backend /go/bin/qanda /usr/bin/qanda
ENTRYPOINT [ "/usr/bin/qanda", "-listen", ":80" ]
EXPOSE 80