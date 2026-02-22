FROM golang:alpine AS builder

WORKDIR $GOPATH/src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY internal ./internal
COPY main.go .

RUN go build -o /go/bin/main
RUN chmod +x /go/bin/main

FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/main /go/bin/main
# Run the hello binary.
ENTRYPOINT ["/go/bin/main"]