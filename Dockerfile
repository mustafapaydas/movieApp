FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git


COPY . /source
WORKDIR /source
ENV DATABASE_NAME=dbname
ENV DATABASE_USERNAME=username
ENV DATABASE_PASSWOR=password
ENV DATABASE_HOST=host
ENV DATABASE_PORT=5432
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV LDFLAGS="-w -s"
# Fetch dependencies.
# Using go get.
RUN go mod tidy
RUN go mod vendor

# Build the binary.
RUN go build -o /source/movie-app

############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /source/movie-app /bin/movie-app
# Run the hello binary.
ENTRYPOINT ["/bin/movie-app"]
