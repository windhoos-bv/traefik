# Windhoos Infra Containers

De container-images die Windhoos intern gebruikt. De images zijn openbaar beschikbaar.

> Dit is niet onze software. De images bouwen voort op bestaande open-sourceprojecten; deze repository bevat alleen onze aanpassingen daarop.

## Traefik

Image: `ghcr.io/windhoos-bv/traefik` (tags `3`, `3.7`, `latest`)

Gebaseerd op de officiële [`traefik`](https://hub.docker.com/_/traefik)-image. Traefik is ontwikkeld door [Traefik Labs](https://traefik.io) en valt onder de [MIT-licentie](https://github.com/traefik/traefik/blob/master/LICENSE.md).

### Wijzigingen

- **CrowdSec-bouncer meegeleverd** – de [crowdsec-bouncer-traefik-plugin](https://github.com/maxlerebourg/crowdsec-bouncer-traefik-plugin) (Apache-2.0) zit als lokale plugin in de image, zodat Traefik hem niet bij het opstarten hoeft te downloaden.
- **Eigen WAF-pagina** – `/waf-forbidden.html` wordt getoond als een verzoek door CrowdSec wordt geblokkeerd.
- **Request-ID-plugin** – eigen plugin (`traefik/3.x/plugins/requestid`) die elk verzoek een willekeurige `X-Request-Id` geeft. Het ID staat in de response-header, de access log en op de WAF-pagina. Een meegestuurde `X-Request-Id` wordt overschreven, tenzij `trustIncoming: true` is gezet.
- **Wekelijkse rebuild** – de image wordt elke week opnieuw gebouwd om beveiligingsupdates van de base-image mee te nemen.

### Configuratie

Static config:

```yaml
experimental:
  localPlugins:
    requestid:
      moduleName: github.com/windhoos-bv/requestid
entryPoints:
  websecure:
    http:
      middlewares: [request-id@file]  # vóór crowdsec, zodat de WAF-pagina het ID kent
accessLog:
  fields:
    headers:
      names:
        X-Request-Id: keep            # verschijnt als request_X-Request-Id
```

Dynamic config:

```yaml
http:
  middlewares:
    request-id:
      plugin:
        requestid: {}
    crowdsec:
      plugin:
        crowdsec-bouncer-traefik-plugin:
          banHTMLFilePath: /waf-forbidden.html
          traceHeadersCustomName: X-Request-Id
```
