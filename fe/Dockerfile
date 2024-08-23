# syntax=docker/dockerfile:1

FROM caddy:2.8.4

ARG UID=10001

COPY .output/public/ /var/sshstats/
COPY Caddyfile /etc/caddy/Caddyfile

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    sshstats

RUN mkdir --parent /var/sshstats/ \
    && \
    chown --recursive sshstats /var/sshstats/

USER sshstats

EXPOSE 65000
