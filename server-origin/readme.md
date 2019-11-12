# Voila-CDN: Origin Server
Origin server of voilaCDN

## Deployment

``` shell

  docker run -d \
  --name vcdn-ori-c1 \
  -v /path/2/configs:/www/voila-CDN-origin/configs \
  -v /path/2/localFileFolder:/www/voila-CDN-origin/statics/files \
  valorad/voila-cdn-origin

  # inspect id
  docker inspect vcdn-ori-c1 | grep '"IPAddress"' | head -n 1

```

## Development

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

Always run the program only when you are at `dist` folder.

``` shell
  # Linux
  $ cd dist
  $ go build -o ./originServer ../src/originServer.go && ./originServer
```

``` powershell
  # Windows
  PS> cd dist
  PS> go build -o .\originServer.exe ..\src\originServer.go && .\originServer.exe
```