# Webfeesh

## Description

Self-hosted RSS collation tool. Single Go Binary. SQLite stored on disk.

For RSS enthusiasts who just want a static HTML page with links to updates made from a collection of feeds.

Given a list of RSS, Atom, or JSON feeds urls, you will be provided a RSS feed link that lists all updates for the collection's feeds since the last poll.

### Features

- RSS, Atom, and JSON feeds are all supported thanks to the [gofeed library](https://github.com/mmcdole/gofeed)
- Web Feed management via Web UI (add, update, delete)
- Background polling per feed with an interval (default: 30 minutes)
- RSS feed link generation with an interval (default: 12 hours)
- New feed item only when an update for a collection has occurred, so your RSS app doesn't get useless updates

## Install

A demo is available [here](). It will refresh every hour.

### Binary

```sh
git clone https://github.com/jangbri/webfeesh.git
# compile vue frontend into frontend/dist
cd webfeesh/frontend
npm install
npm run build
# build the single binary
cd ..
go mod download
go build cmd/web/main.go
```

### Docker

```sh
git clone https://github.com/jangbri/webfeesh.git
cd webfeesh
docker compose build
docker compose up -d
```

## Dev Setup

Pre-requisites:

- npm
- Golang
- Sqlite

This project is primarily built in Go and Vue 3. The steps provided for building your own binary is sufficient for also setting up a dev environment.
