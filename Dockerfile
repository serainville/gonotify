FROM golang:1.17.2

COPY . /go/src/gonotify

WORKDIR /go/src/gonotify

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /gonotify

FROM scratch

COPY --from=0 /gonotify /gonotify

ENTRYPOINT [ "/gonotify" ]