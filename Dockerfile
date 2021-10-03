FROM golang:1.17.1 AS buildenv
WORKDIR /build/src
COPY . .
COPY .git .

RUN make BUILD_OUTPUT=/build/out release
RUN make release-test
RUN cp /build/out/server /build/root/

RUN mkdir -p /build/root/tmp
RUN chmod 1777 /build/root/tmp

FROM scratch
WORKDIR /
COPY --from=buildenv /build/root/ /

USER 10000:10000
EXPOSE 8080/tcp
EXPOSE 9090/tcp
EXPOSE 12345/tcp
EXPOSE 6060/tcp
ENV TMPDIR=/tmp

CMD ["/server", \
    "-rest_port", "8080", \
    "-monitoring_port", "9090", \
    "-grpc_port", "12345", \
    "-pprof_port", "6060"]
