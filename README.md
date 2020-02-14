# Ferro
If you don't already know how to use it, you probably shouldn't be. If you want to anyways, check out `./docs/` and ask `Zero#5200` on discord for a secret key

### TODO
- [x] POST /
- [ ] GET /:id/
    - [x] file
    - [ ] ?cropx
    - [ ] ?cropxy
    - [ ] ?scale
    - [ ] ?scalex
    - [ ] ?scaley
- [x] DELETE /:id/
- [ ] GET /:id/md5
- [ ] GET /:id/thumb

### Environment Variables

MONGO_USER -> Mongodb username
MONGO_PASS -> Mongodb password
MONGO_HOST -> Mongodb server host
FERRO_SECRET -> Auth header value to look for
