FROM debian:stable
LABEL maintainer="raghwendrasingh1819@gmail.com"

# Extract application tarball
RUN mkdir -p /go/bin
ADD build/output/go-endpoint /go/bin
WORKDIR /go/bin

COPY build/runtime/entrypoint.sh /
CMD [ "bash",  "/entrypoint.sh"  ]
