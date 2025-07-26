# ddnsman

[![codecov](https://codecov.io/github/leonidboykov/ddnsman/graph/badge.svg?token=FUdTxriSvs)](https://codecov.io/github/leonidboykov/ddnsman)

ddnsman (short for Dynamic DNS Manager) allows to set external IP as a DNS record for your DNS provider.

## Supported providers

ddnsman uses [`github.com/libdns`](https://github.com/libdns) to communicate with a list of various providers, i.e. if
`libdns` allows to work with specific DNS provider you may use it with `ddnsman`. Several providers does not support
requred methods, thus they may be absent.

## Settings

Here is an example:

``` jsonc
{
  "interval": "5m", // See https://pkg.go.dev/time#ParseDuration.
  "settings": [
    {
      "domain": "example.com", // Your domain.
      "provider": {
        "name": "someprovider", // Any provider supported by libdns.
        "settings": {} // This data is passed directly to provider driver.
      },
      "records": [ // Any records you want to update.
        "subdomain",
        "*.wildcard",
      ]
    }
  ],
  "shoutrrr_notifications": [
    {
      "url": "...", // shoutrrr url.
      "settings": {} // additional settings for shoutrrr (just a QoL feature, you may use params as well).
    }
  ]
}
```

Proper docs are coming soon.
