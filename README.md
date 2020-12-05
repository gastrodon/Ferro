# Ferrothorn

### Getting Started

To get a build up, you can use `docker-compose up`.
This will create a running instance of Ferro in a container exposed at `localhost:80` and a mongo docker image for it to talk to.
Relevant files are stored at `/ferrothorn`.

When starting the regular stack (`docker-compose up`), the same thing happens but with files being stored at `/ferrothorn`.

### How to use

More information can be found inside of the `docs` folder of this repository. Docs can be built with [persim](https://github.com/gastrodon/persim) locally.

In short, you can `POST`, `GET`, and `DELETE` arbitrary content to ferrothorn. If `FERROTHORN_SECRET` is set, any request that changes content (ie not a `GET`) will require a header `Authentication: $FERROTHORN_SECRET`.

This is intended as a fileserve for a single application (ie some application will have it's own ferrothorn), so multiple secrets or a login / logout API is not planned.

### Environment Variables
```
FERROTHORN_ROOT         -> Root directory for all files
FERROTHORN_CONNECTION   -> mysql connection string
FERROTHORN_SECRET       -> Authorization header needed to make changes
```
