FROM --platform=${BUILDPLATFORM:-linux/amd64} golang as builder

ARG TARGETOS
ARG TARGETARCH

ENV CGO_ENABLED=0

COPY . .

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o ddnsman ./cmd/ddnsman

FROM --platform=${BUILDPLATFORM:-linux/amd64} alpine

COPY --from=builder ddnsman /usr/local/bin
WORKDIR /
ENTRYPOINT ["/bin/sh", "-c"]
CMD ["ddnsman"]
