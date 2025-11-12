[![Test](https://github.com/dyammarcano/j2t/actions/workflows/test.yml/badge.svg)](https://github.com/dyammarcano/j2t/actions/workflows/test.yml)

j2t — JSON → Toon encoder (encode-only)

j2t is a small command-line tool that encodes JSON into the "Toon" format using the `gotoon` library.

The CLI exposes a single subcommand, `encode`. It reads JSON from a file or from stdin and writes Toon to stdout by default. Use `-o` / `--output` to write to a file.

Requirements
------------
- Go (a recent stable version). Ensure your `GOPATH/bin` is on `PATH` to run installed binaries.

Install
-------
You can install the binary into your Go bin directory in two common ways:

- Directly from the module path:

```
go install github.com/dyammarcano/j2t@latest
```

Usage
-----

Encode a JSON file and print the Toon output to stdout:

```
j2t encode input.json
```

Encode JSON from stdin:

```
cat input.json | j2t encode
# Windows (cmd.exe) alternative:
type input.json | j2t encode
```

Write encoded output to a file:

```
j2t encode input.json -o output.toon

# or from stdin
cat input.json | j2t encode -o output.toon
# Windows (cmd.exe) alternative:
type input.json | j2t encode -o output.toon
```

Show help:

```
j2t --help
j2t encode --help
```

Flags
-----
- `-o, --output`  write output to the specified file (default: stdout)
- `-h, --help`    show help for commands

Notes
-----
- The CLI currently implements encoding (JSON → Toon). Decoding was intentionally removed.
- The project depends on `github.com/alpkeskin/gotoon` for encoding and `github.com/spf13/cobra` for CLI scaffolding. Dependencies are managed in `go.mod`.

License
-------
This project is released under the MIT License. See the `LICENSE` file for details.

Reference
---------
- Toon format specification and implementation: https://github.com/toon-format/toon — the canonical repo for the Toon format used by this tool.
