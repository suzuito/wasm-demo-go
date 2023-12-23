# WASM demo app

This is demo app of WASM in Go.

It is assumed that this demo is run on Mac.

## Build

```bash
DEMO=demo001 make
```

Hot reload wanted? Use [air](https://github.com/cosmtrek/air)!

```bash
# If air is not installed,
# go install github.com/cosmtrek/air@latest
DEMO=demo001 $(go env GOPATH)/bin/air -c .air.toml
```

## Run

```bash
python3 -m http.server -d docs/
```
