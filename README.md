# tsd
timestamp to date ISO format

## Build and run
```
$ go build -o tsd ./cmd/tsd

$ ./tsd 1833561154
2024-12-07 08:45:54 +0000 UTC
```

## Test
```
$ gotest -v -cover ./...  # install gotest
```

## Install from go.pkg.dev
```
$ go install github.com/nellaG/tsd/cmd/tsd@{version}
```
`
