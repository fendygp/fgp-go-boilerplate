# Step 1 build executable binary
FROM golang:alpine as builder

# Install git
COPY . $GOPATH/src/fgp-go-boilerplate/
WORKDIR $GOPATH/src/fgp-go-boilerplate/

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

#build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /fgp-go-boilerplate .


FROM alpine:3.4

# Copy our static executable
COPY --from=builder /fgp-go-boilerplate /fgp-go-boilerplate
ENTRYPOINT ["./fgp-go-boilerplate"]