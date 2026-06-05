# concord-plugin-hello

Reference Concord plugin. Demonstrates the v1 plugin protocol end-to-end:
a single `greeting` evidence type and a `Probe` health check.

## Build

```sh
make build
```

Produces `bin/concord-plugin-hello`.

## Install locally

```sh
make install        # copies the binary to ~/.concord/plugins/hello/v0.1.0/
concord check --controls ./testdata
```

## Author a similar plugin

Copy this repo as a template, change `hello` to your source name, and
replace the `Collect` implementation. See `docs/writing-a-plugin.md` in
the concord repo for the full guide.

## Protocol version

This plugin speaks the v1 Concord plugin protocol. Wire definition lives
at `github.com/concord-dev/concord/proto/concord/plugin/v1/collector.proto`.
