# Image that provides cross compilation tooling.
FROM --platform=$BUILDPLATFORM rancher/mirrored-tonistiigi-xx:1.6.1 AS xx

# Minimal BCI base image, as recommended by RFD 0070.
FROM registry.suse.com/bci/bci-nano:16.0 AS base

FROM --platform=$BUILDPLATFORM registry.suse.com/bci/golang:1.27 AS builder

# bci-nano has no shell to run commands in, so its identity files are
# prepared here. Unlike BCI, it does not define the nobody user/group,
# which the final image runs as.
COPY --from=base /etc/group /final/etc/group
RUN echo "nobody:x:65534:" >> /final/etc/group && \
    install -dv -m 0755 /final/var/lib/nobody && \
    cp /etc/nsswitch.conf /final/etc

# Differs from BCI, by removing /bin/bash from root:
RUN printf "root:x:0:0:root:/root:/usr/bin/false\nnobody:x:65534:65534:nobody:/var/lib/nobody:/usr/bin/false\n" > /final/etc/passwd

WORKDIR /src
COPY go.sum \
     go.mod \
     Makefile \
     /src
RUN go mod download

COPY pkg /src/pkg
COPY hack /src/hack
COPY main.go /src

# Ensures that the binary that was built was cross-compiled correctly
# and is valid on the target platform.
COPY --from=xx / /

ARG TARGETPLATFORM
RUN xx-go --wrap && mkdir -p /run/lock

# By setting the version as an argument, we can avoid running the version logic 
# a second time (inside the Docker build process). Therefore, removing the need
# to access the .git dir.
ARG VERSION
RUN VERSION=${VERSION} TARGET_BIN=/out/compliance-operator make build
RUN xx-verify --static /out/compliance-operator

FROM base AS final

COPY --from=builder /final/ /
COPY --from=builder /out/compliance-operator /usr/bin/compliance-operator

# Aligns nobody user ID with BCI.
USER 65534:65534

CMD ["/usr/bin/compliance-operator"]
