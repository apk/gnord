# A go http server

Modeled after some features of [fnord](https://www.fefe.de/fnord/),
a C webserver.

## Features

Allows good old cgi scripts (and only those).
They must have the suffix `.cgi`,
but are referenced without the suffix, so `/test` will execute
`test.cgi`, and `/test.cgi` will return a 404. This is explicitly
so the cgi source won't be readable. (Any URL ending `.cgi` will
yield a 404.)

Allows symlinks for redirects. `ln -s https://google.com/ g`
makes `/g` redirect to google.

Using the content type detection of golang http, so you can
easily have page URLs without the `.html` ending.

## Non-features

Only serves a single document tree - no virtual hosts.
You are expected to run this behind nginx or similar,
even if just for SSL. (Or for a [tor](https://www.torproject.org/)
onion service.)

## Options

By default it serves the current directory
and thus must be started there. Use
`--path <dir>` to change.

Default listener port is `127.0.0.1:4040`.
Use `--addr host:port` to change. Binding
to localhost is recommended, but not enforced
(nor made easy in any way).

If `--ip <name>` is given the named request header
(if present) will be used instead of the connection
remote IP for `REMOTE_ADDR` in cgi scripts. E.g.
`--ip X-Forwarded-For`. Security considerations
apply - your proxy should unconditionally set the
header you specify here.

`--well-known <host>` allows to forward requests
to `/.well-known` (let's encrypt validator) to
`https://host/.well-known`, to allow to do the
actual cert generation elsewhere.

## Bugs

No SSL yet.

Directory listings still contain the cgi scripts
under their actual `t.cgi` name, not as `t`. Directory
listings could be better anyway.
