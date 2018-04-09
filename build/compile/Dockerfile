FROM golang:1.9
LABEL maintainer="raghwendrasingh1819@gmail.com"

# Setup output directory for application binary
VOLUME /output
WORKDIR /output

# Update entrypoint to build and package application
COPY entrypoint.sh /
CMD [ "bash",  "/entrypoint.sh"  ]
