# syntax=docker/dockerfile:1

ARG GO_VERSION=1.22.1

FROM golang:${GO_VERSION} AS build

ENV CGO_ENABLED=0
ENV GOARCH=amd64

WORKDIR /src

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    go build -o /bin/server ./cmd/api

FROM alpine:latest AS final

RUN --mount=type=cache,target=/var/cache/apk \
    apk --update add \
    ca-certificates \
    tzdata \
    && \
    update-ca-certificates

ARG UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser

RUN mkdir /etc/sshstats 

USER appuser

COPY --from=build /bin/server /bin/

EXPOSE 65535

ENV GIN_MODE=release

ENV SSHSTATS_SERVER_ADDRESS=:65535
ENV SSHSTATS_CONFIG_DIR=/etc/sshstats

ENV SSHSTATS_AUTH_HEADERKEY=
ENV SSHSTATS_AUTH_HEADERVALUE=

ENV SSHSTATS_LOKI_ENDPOINT=
ENV SSHSTATS_LOKI_USERNAME=
ENV SSHSTATS_LOKI_PASSWORD=

ENTRYPOINT [ "/bin/server" ]
