FROM alpine

COPY ddnsman /usr/local/bin
WORKDIR /

ENTRYPOINT ["/bin/sh", "-c"]
CMD ["ddnsman"]
