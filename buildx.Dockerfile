FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:alpine as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o ddnsman

FROM --platform=${BUILDPLATFORM:-linux/amd64} alpine

COPY --from=builder ddnsman /usr/local/bin
WORKDIR /
ENTRYPOINT ["/bin/sh", "-c"]
CMD ["ddnsman"]
