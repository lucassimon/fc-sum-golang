FROM golang:1.13-alpine AS build

COPY /build/sum /bin/sum
ENTRYPOINT ["/bin/sum"]
