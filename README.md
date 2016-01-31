# Site Heartbeat

Checks for a bunch of domains their `<title>`.

Uses:
- [Goquery](https://github.com/PuerkitoBio/goquery)

## Install deps

```
go get github.com/PuerkitoBio/goquery
```

## Build

```
go build siteheartbeat
```

## Configs

`hosts.txt`: List of domains and their title (regexp) -> see `hosts.txt.dist` as template

`mailer.conf`: Mailer config to send the "Site down" mails (yet not configurable)

## Run

```
./siteheartbeat
```

## Disclaimer and Contribute

It's a prototype and hardly situated for my use case. 

If you want to have more configurations, more mailer setups, more flags, I'm curious to see your PR :) So fork it and try your best. It's just Go.
