# Voila-CDN: Origin Server

## dev
Fetch package

``` shell
  $ go build
  $ go mod tidy
  $ go mod vendor
```

Run
``` shell
  $ go build -o dist/originServer src/originServer.go && dist/originServer
```

``` powershell
  PS> go build -o dist/originServer.exe src/originServer.go && dist/originServer.exe
```

