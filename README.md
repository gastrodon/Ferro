# Ferrothorn

### Getting Started

Ferro requires the following to run
- `Go 1.13+`
- `Mongo 4.2+`

If building with docker, Ferro requires
- `docker`
- `docker-compose`

To get a build up, you can use `docker-compose up`.
This will create a running instance of Ferro in a container expoed at `localhost:8000` and a mongo docker image for it to talk to.
Relevant files are stored at `/ferrothorn`.

When starting the regular stack (`docker-compose up`), the same thing happens but with files being stored at `/ferrothorn`.

### How to use

More information can be found inside of the `docs` folder of this repository. Docs can be built with [persim](https://github.com/gastrodon/persim) locally.

In short, you can `POST`, `GET`, and `DELETE` arbitrary content to ferrothorn. If `FERRO_SECRET` is set, any request that changes content (ie not a `GET`) will require a header `Authentication: $FERRO_SECRET`.

This is intended as a fileserve for a single application (ie some application will have it's own ferrothorn), so multiple secrets or a login / logout interface is planned to be supported.

### Environment Variables
```
FERRO_MONGO_USER -> Mongodb username
FERRO_MONGO_PASS -> Mongodb password
FERRO_MONGO_HOST -> Mongodb server host
FERRO_MONGO_BASE -> Mongodb database to talk to
FERRO_SECRET     -> Auth header value to look for
FERRO_LOG_LEVEL  -> Logger severity level, 2 is most verbose
```

### TODO
- [x] POST /
- [x] GET /:id/
    - [x] file
- [x] DELETE /:id/
- [ ] GET /:id/md5
