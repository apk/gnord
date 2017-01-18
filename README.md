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

It serves the current directory and thus must be
started there.

## Other

Default listener port ist `127.0.0.1:4040`.
