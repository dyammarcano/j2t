j2t — JSON ↔ Toon encoder (encode-only)

j2t is a small command-line tool that encodes JSON into the "Toon" format using the `gotoon` library.

This repository provides a simple CLI with the `encode` command. It reads from a file or stdin and writes to stdout by default. Use `-o` / `--output` to save the result to a file.

Build
-----
From the repository root (Windows cmd.exe):

```
:: Build a local binary named j2t
go build -o j2t .
```

Install
-------
You can install the binary into your Go bin directory in two common ways:

- Install from the local module (installs the package in the current directory):

```
:: Install the binary from the current module into $(go env GOPATH)\bin
go install .
```

- Install from a remote module path (use this when the module has a real import path, replace `<module/path>`):

```
:: Install the latest released version from the module proxy
go install <module/path>@latest
```

Note: replace `<module/path>` with the actual module import path (the `module` line in `go.mod`) if you publish this repository. If you're working locally `go install .` is the simplest option.

You can also install a binary into your GOPATH/bin (if module path is set correctly):

```
go install
```

Usage
-----

Encode a JSON file and print the Toon output to stdout:

```
j2t encode input.json
```

Encode JSON from stdin (Windows cmd.exe uses `type` to stream a file):

```
type input.json | j2t encode
```

Write encoded output to a file:

```
j2t encode input.json -o output.toon

:: or from stdin
type input.json | j2t encode -o output.toon
```

Flags
-----
- `-o, --output`  write output to the specified file (default: stdout)
- `-h, --help`    show help for commands

Notes
-----
- The CLI currently implements encoding (JSON -> Toon). Decoding was intentionally removed.
- The project depends on `github.com/alpkeskin/gotoon` for encoding and `github.com/spf13/cobra` for CLI scaffolding. Dependencies are managed in `go.mod`.

License
-------
This project is released under the MIT License. See the `LICENSE` file for details.

Reference
---------
- Toon format specification and implementation: https://github.com/toon-format/toon — the canonical repo for the Toon format used by this tool.
