FROM golang:alpine AS builder

WORKDIR $GOPATH/src
COPY src .
# Fetch dependencies.
# Using go get.
# Build the binary.
RUN go build -o /go/bin/main
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/main /go/bin/main
# Run the hello binary.
ENTRYPOINT ["/go/bin/main"]