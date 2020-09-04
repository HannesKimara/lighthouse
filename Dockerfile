FROM golang AS build

WORKDIR /go/src/app
COPY . /go/src/app

RUN go get -d -v && \
    go build -o /go/bin/app


# Copy into base image
FROM gcr.io/distroless/base

WORKDIR /opt

COPY --from=build /go/bin/app /opt/
COPY --from=build /go/src/app/conf.yml /opt/

CMD ["/opt/app"]
