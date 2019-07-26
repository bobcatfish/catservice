FROM golang:1.12 as build

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

FROM gcr.io/distroless/base
COPY --from=build /go/src/app/static /static
COPY --from=build /go/bin/app /
CMD ["/app"]
