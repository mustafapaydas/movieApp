FROM golang:alpine AS builder

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

RUN go mod tidy
RUN go mod vendor
RUN go build -o /source/movie-app


FROM scratch
COPY --from=builder /source/movie-app /bin/movie-app
ENTRYPOINT ["/bin/movie-app"]
