# Site Heartbeat

Checks for a bunch of domains their `<title>`.

Uses:
- [Goquery](https://github.com/PuerkitoBio/goquery)

## depends on Go 1.5

On Ubuntu 14.04 the standard is Go 1.2. To install Go 1.5 [follow this instructions](http://munchpress.com/install-golang-1-5-on-ubuntu/).

## Install

```
go get github.com/PuerkitoBio/goquery
go get github.com/ronnyhartenstein/site-heartbeat-golang
```

## Build

```
go build github.com/ronnyhartenstein/site-heartbeat-golang
```

Build for other target platform:

```
env GOOS=linux GOARCH=386 go build github.com/ronnyhartenstein/site-heartbeat-golang
```

## Configs

`hosts.txt`: List of domains and their title (regexp) -> see `hosts.txt.dist` as template

`mailer.conf`: Mailer config to send the "Site down" mails (yet not configurable) -> see `mailer.conf.dist` as template

## Run

```
./siteheartbeat
```

## Disclaimer and Contribute

It's a prototype and hardly situated for my use case.

If you want to have more configurations, more mailer setups, more flags, I'm curious to see your PR :) So fork it and try your best. It's just Go.
