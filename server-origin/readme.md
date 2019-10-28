# Voila-CDN: Origin Server

## dev

Fetch package

``` shell
  $ go build
  $ go mod tidy
  $ go mod vendor
```

Get the dist folder ready

- Create the `dist` folder
- then create a `voila-CDN-origin.json` under `dist/configs` folder
- put the files you like under `statics` folder

Run

Always run the program when you are at `dist` folder.

``` shell
  # Linux
  $ cd dist
  $ go build -o ./originServer ../src/originServer.go && originServer
```

``` powershell
  # Windows
  PS> cd dist
  PS> go build -o .\originServer.exe ..\src\originServer.go && .\originServer.exe
```

