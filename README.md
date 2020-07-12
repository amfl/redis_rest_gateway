# Redis REST Gateway

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
