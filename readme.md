# Windhoos Infra Containers

Deze repository bevat herbruikbare containers die we binnen (en buiten) Windhoos gebruiken.

Let op dat alle containers in deze repository public zijn.

## Traefik

### Request-ID

De image bevat een lokale plugin (`traefik/3.x/plugins/requestid`) die elk verzoek een willekeurige `X-Request-Id` geeft. Een meegestuurde `X-Request-Id` wordt standaard overschreven; zet `trustIncoming: true` om die te behouden. Het ID staat in de response-header, in de access log en op de WAF-pagina.

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
