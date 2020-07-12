# Redis REST Gateway

[![Docker Automated Build](https://img.shields.io/docker/cloud/automated/amfl/redis_rest_gateway)](https://hub.docker.com/r/amfl/redis_rest_gateway)
[![Go Report Card](https://goreportcard.com/badge/github.com/amfl/redis_rest_gateway)](https://goreportcard.com/report/github.com/amfl/redis_rest_gateway)

Webserver which passes POST data and headers to Redis.

Originally made as a way to push Gitea events into Redis,
but can be used generically.

There is **no authentication**.

## Usage

```bash
rrgw LISTEN_INTERFACE REDIS_ADDR
```

For example, to listen on 8080 on all interfaces and publish to the redis at
localhost:6379, use:

```bash
rrgw :8080 localhost:6379
```

You can then test it with:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"foo": "bar"}' localhost:8080/channel/foo
```

## Similar Projects

[Webdis](https://webd.is/) solves the same problem, but it is either not
flexible enough to pass along headers and accept arbitrary JSON POST data, or I
am not smart enough to figure it out :)
