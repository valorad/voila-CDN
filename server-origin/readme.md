# Voila-CDN: Origin Server

Origin server of voila-CDN

Check out the [API document on Apiary](https://voilacdnorigin.docs.apiary.io)

## Deployment

- Create a folder for storing files, for later to upload to replica server. 
- Create a config file called `voila-CDN-origin.json`, and put it under the folder `configs`
- Configure the `replicaHosts`, which is the list of addresses of replica servers.
- run the following command. You may also specify the port to map (e.g. add `-p 13399:3399` after `docker run ...`)

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

### Fetch package

``` shell
  $ go build
  $ go mod tidy
  $ go mod vendor
```

### Get the dist folder ready

- Create the `dist` folder
- then create a `voila-CDN-origin.json` under `dist/configs` folder
- put the files you like under `statics/files` folder

### Run

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