# Server Replica
Replica server

## Deployment

  A django `settings.py` is needed before running the container.

    ``` shell

    docker run -d \
      --name vcdn-rep-c1 \
      -v /path/2/media:/www/voila-CDN-replica/media \
      -v /path/2/settings.py:/www/voila-CDN-replica/replicaServer/settings.py \
      valorad/voila-cdn-replica
      
    ```