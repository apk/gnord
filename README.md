# A go http server

Modeled after [https://www.fefe.de/fnord/](fnord), a C webserver.

## Features

Allows good old cgi scripts. Those must have the suffix `.cgi`,
but are referenced without the suffix, so `/test` will execute
`test.cgi`, and `/test.cgi` will return a 404.

Allows symlinks for redirects. `ln -s https://google.com/ g`
makes `/g` redirect to google.

## Non-features

Only serves a single document tree - no virtual hosts.
You are expected to run this behind nginx or similar,
even if just for SSL.


## Other

By default it serves the current directory
and thus must be started there. Use
`--path <dir>` to change.

Default listener port is `127.0.0.1:4040`.
Use `--addr host:port` to change. Binding
to localhost is recommended, but not enforced
(nor made easy in any way).
