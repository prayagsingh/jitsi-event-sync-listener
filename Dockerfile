# Using golang:1.17.7-alpine3.15
FROM golang@sha256:1dc6a836407ef26c761af27bd39eb86ec385bab0f89a6c969bb1a04b342f7074 AS builder

# Install git + SSL ca certificates(for exposing this app on https).
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create appuser.
ENV USER=appuser
ENV UID=10001 

# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR $GOPATH/src/github.com/prayagsingh/jitsi-event-sync-listner/

COPY . .

# Fetch dependencies.
RUN go mod download
RUN go mod verify

# Build the binary.
# using CGO_ENABLED=0 because https://stackoverflow.com/questions/55106186/no-such-file-or-directory-with-docker-scratch-image
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/jitsi-event-sync-listner

############################
# STEP 2 build a small image
############################
FROM scratch

# Import the user and group files from the builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable.
COPY --from=builder /go/bin/jitsi-event-sync-listner /go/bin/jitsi-event-sync-listner

# Use an unprivileged user.
USER appuser:appuser

# Port on which the service will be exposed.
EXPOSE 7002

# Run the binary.
ENTRYPOINT ["/go/bin/jitsi-event-sync-listner"]