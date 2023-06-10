FROM --platform=${BUILDPLATFORM:-linux/amd64} golang as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

COPY . .

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o ddnsman

FROM --platform=${BUILDPLATFORM:-linux/amd64} alpine

COPY --from=builder ddnsman /usr/local/bin
WORKDIR /
ENTRYPOINT ["/bin/sh", "-c"]
CMD ["ddnsman"]
