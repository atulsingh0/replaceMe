FROM golang:1.19 as build
WORKDIR /build
COPY cmd/ .
RUN go get
RUN CGO_ENABLED=0 go build

# Unprivileged users can execute
RUN chgrp 0 replaceme
RUN chmod g+x replaceme

FROM scratch
COPY --from=build /build/replaceme .
USER 65534
ENTRYPOINT ["/replaceme"]
