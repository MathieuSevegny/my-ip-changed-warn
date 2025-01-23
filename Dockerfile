FROM golang:alpine AS builder

WORKDIR $GOPATH/src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY internal ./internal
COPY main.go .

# Fetch dependencies.
# Using go get.
# Build the binary.
RUN go build -o /go/bin/main
RUN chmod +x /go/bin/main

############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/main /go/bin/main
ENV API_ENDPOINT=https://api.ipify.org
ENV CACHE_FOLDER_PATH=/cache
ENV CACHE_FILENAME=last_ip.txt
ENV EMAIL_TO=test@gmail.com
ENV EMAIL_FROM=test@gmail.com
ENV SECONDS_TO_WAIT=10
ENV DEVICE_NAME=YOU_CAN_CHANGE_THIS_NAME
ENV MAX_TRIES=10
# Run the hello binary.
ENTRYPOINT ["/go/bin/main"]