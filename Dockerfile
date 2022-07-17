FROM golang:1.16 as build
WORKDIR /build
COPY src/ .
RUN go get
RUN CGO_ENABLED=0 go build

# Unprivileged users can execute
RUN chgrp 0 replaceMe
RUN chmod g+x replaceMe

FROM scratch
COPY --from=build /build/replaceMe .
USER 65534
ENTRYPOINT ["/replaceMe"]
