# syntax = docker/dockerfile:1-experimental
# instantiate and define global build args
ARG exec=points-exercise
ARG port=8080

FROM golang:1.15 AS base
# ensure args are referenced
ENV CGO_ENABLED=0
WORKDIR /src
# add all development files into relative destination
# use dockerignore which ignores tests and misc files
# copy go.mod and go.sum
COPY ./go.mod .
COPY ./go.sum .
# download dependencies if go.* changes
RUN go mod download

FROM base AS build
ARG exec
RUN --mount=target=. \
--mount=type=cache,target=/root/tmp/go-build \
CGO_ENABLED=0 GOOS=linux go build -o /out/${exec} .

FROM base AS unit-test
RUN --mount=target=. \
--mount=type=cache,target=/root/tmp/go-build \
go test -v .

FROM golangci/golangci-lint:v1.42-alpine AS lint-base

FROM base AS lint
ARG exec
WORKDIR /app
RUN --mount=target=. \
--mount=from=lint-base,src=/usr/bin/golangci-lint,target=/usr/bin/golangci-lint \
--mount=from=cache,target=/root/tmp/go-build \
--mount=type=cache,target=/root/tmp/golangci-lint \
CGO_ENABLED=0 GOOS=linux go build -o /out/${exec} .

# copy resources
FROM alpine:latest AS deployment
RUN apk --no-cache add ca-certificates
# ensure args are referenced
ARG exec
WORKDIR /app
# add necessary html files
COPY ./html/ ./html/
# copy binary from built image into working directory
COPY --from=build /out/${exec} .


# for localhost development
FROM deployment AS localhost
ARG port
LABEL description="points exercise with golang and docker" \
      maintainer="phil"
WORKDIR /app
EXPOSE ${port}
CMD ["./points-exercise"]
