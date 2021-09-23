FROM golang:1.17.1 AS buildenv
WORKDIR /build/src
COPY go.mod .
COPY go.sum .
ADD api api
ADD cmd cmd
ADD pkg pkg

RUN mkdir -p /build/out
RUN CGO_ENABLED=0 go build -ldflags "-w" -o /build/out ./...

FROM gcr.io/distroless/base
COPY --from=buildenv /build/out/server /

USER 1000:1000
EXPOSE 8080/tcp
EXPOSE 8090/tcp
CMD ["/server", "-grpc_port", "8080", "-rest_port", "8090"]
