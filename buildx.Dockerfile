FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:alpine as builder

ARG TARGETOS
ARG TARGETARCH
WORKDIR /src

ENV CGO_ENABLED=0

COPY . .

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o ddnsman ./cmd/ddnsman

FROM alpine

COPY --from=builder /src/ddnsman /usr/local/bin
WORKDIR /
ENTRYPOINT ["/bin/sh", "-c"]
CMD ["ddnsman"]
