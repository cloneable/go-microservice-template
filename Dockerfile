FROM golang:1.17.1 AS buildenv
WORKDIR /build/src
COPY . .
COPY .git .

RUN make BUILD_OUTPUT=/build/out release
RUN make release-test

RUN mkdir -p /build/root
RUN cp /build/out/server /build/root/
RUN mkdir /build/root/tmp
RUN chmod 1777 /build/root/tmp

FROM scratch
WORKDIR /
COPY --from=buildenv /build/root/ /

USER 10000:10000
EXPOSE 12345/tcp
EXPOSE 8080/tcp
EXPOSE 9090/tcp
ENV TMPDIR=/tmp

CMD ["/server", \
    "-port", "12345", \
    "-rest_port", "8080", \
    "-grpc_port", "9090"]
