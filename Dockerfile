# create a base stage for basic requirements
FROM golang:1.25 AS build-base

WORKDIR /jamitize-server

# copy the go dependencies and libraries
COPY go.mod go.sum ./

# caches the required to installations to avoid installing again upon change (saves a lot of time downloading)
RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go mod download

# create a new stage for development build
FROM build-base AS build-dev

# enables hot reload for frequent changes reflection (only for dev environment)
RUN go install github.com/cosmtrek/air@v1.43.0 && \
  go install github.com/go-delve/delve/cmd/dlv@v1.20.2

# copy the source code into workdir
COPY . .

# enable hot reload
CMD ["air", "-c", ".air.toml"]

# create a new stage for production build
FROM build-base AS build-prod

# create a non root user
RUN useradd -u 1001 nonroot

# copy the source code into workdir
COPY . .

# builds a binary at the compile time to check for any errors
RUN go build \
  -ldflags="-linkmode external -extldflags -static" \
  -tags netgo \
  -o bin/api \
  ./cmd/api/main.go

# build a binary for health check to see if the application server works or not
RUN go build \
  -ldflags="-linkmode external -extldflags -static" \
  -tags netgo \
  -o bin/healthcheck \
  ./healthcheck/healthcheck.go

# use another stage for deployable image
FROM scratch

ENV GIN_MODE=RELEASE

# Copy the password file for user to access in production
COPY --from=build-prod /etc/passwd /etc/passwd

# Copy the app binary from the build stage into the scratch stage
COPY --from=build-prod /jamitize-server/bin/api bin/api

# Use non root user
USER nonroot

# only for apis
EXPOSE 8080

# run binary
CMD ["/bin/api"]
