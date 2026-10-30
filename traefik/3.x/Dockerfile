FROM alpine:3 AS crowdsec-plugin
ARG CROWDSEC_PLUGIN_VERSION=v1.8.1
RUN apk add --no-cache git \
 && git clone --depth 1 --branch ${CROWDSEC_PLUGIN_VERSION} \
    https://github.com/maxlerebourg/crowdsec-bouncer-traefik-plugin.git /plugin \
 && rm -rf /plugin/.git

FROM traefik:v3.7

# Load via experimental.localPlugins, so Traefik does not download the plugin on startup
COPY --from=crowdsec-plugin /plugin /plugins-local/src/github.com/maxlerebourg/crowdsec-bouncer-traefik-plugin
COPY waf-forbidden.html /waf-forbidden.html
