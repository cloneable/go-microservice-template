FROM golang:1.17.1 AS buildenv
WORKDIR /build/src
COPY . .
COPY .git .

ENV CGO_ENABLED=0
ENV GO111MODULE=on
RUN mkdir -p /build/out
RUN go build \
    -trimpath \
    -ldflags "-s -w" \
    -installsuffix cgo \
    -o /build/out \
    ./...
RUN go test \
    -v \
    -trimpath \
    -ldflags "-s -w" \
    -installsuffix cgo \
    ./...

FROM scratch
COPY --from=buildenv /build/out/server /

USER 10000:10000
EXPOSE 8080/tcp
EXPOSE 9090/tcp
EXPOSE 12345/tcp
EXPOSE 6060/tcp

CMD ["/server", "-rest_port", "8080", "-monitoring_port", "9090", "-grpc_port", "12345", "-pprof_port", "6060"]
