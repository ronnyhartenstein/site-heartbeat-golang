# Site Heartbeat

Checks for a bunch of domains their HTTP status code (200) and the `<title>`.

Uses:
- [Goquery](https://github.com/PuerkitoBio/goquery)

## Install deps

```
go get github.com/PuerkitoBio/goquery
```

## Run

```
go build siteheartbeat && ./siteheartbeat
```
