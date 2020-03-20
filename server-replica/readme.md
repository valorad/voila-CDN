# Server Replica

[![Docker pulls](https://img.shields.io/docker/pulls/valorad/voila-cdn-replica.svg?style=flat-square)](https://hub.docker.com/r/valorad/voila-cdn-replica/)
[![](https://images.microbadger.com/badges/version/valorad/voila-cdn-replica.svg)](https://microbadger.com/images/valorad/voila-cdn-replica "voila-cdn-replica Version")
[![](https://images.microbadger.com/badges/image/valorad/voila-cdn-replica.svg)](https://microbadger.com/images/valorad/voila-cdn-replica "voila-cdn-replica Image")

Replica server of voila-CDN.

Check out the [API document on Apiary](https://voilacdnreplica.docs.apiary.io)

## Deployment

  ### Via `docker-compose`:

  - Create your own `docker-compose.yml` by following the example file.
  - Specify the `volumes` and `networks.voila-cdn_net.ipv4_address` fields according to your needs.
  - For EACH replica server, create a folder
  - Copy the `settings.example.py` from `./replicaServer/` to EACH replica server folder, and rename to `settings.py`
  - Specify the secret keys at `SECRET_KEY = ` field in each `settings.py`. The [secret key](https://docs.djangoproject.com/en/dev/ref/settings/#secret-key) is a random string that consists of 50 characters ranging from `[a-z][0-9][!@#$%^&*(-_=+)]`.
    - You may use `django startproject`, or some 3rd party tools like [Strong Random Password Generator](https://passwordsgenerator.net/) to genrate the secret keys.
  - For each `settings.py`, at `ALLOWED_HOSTS = []` list, add the `ipv4_address` value you specified for this service earlier in `docker-compose.yml`. 
    - Example: `ALLOWED_HOSTS = ["172.22.0.100"]` for `replica1`, `ALLOWED_HOSTS = ["172.22.0.101"]` for `replica2`, etc.
  - create a `media` folder for each replica server, which will be used to store files sent from the Origin server
  - run

    ``` shell
      $ docker-compose up -d
    ```

  ### On Kubernetes:

  Prepare your single `settings.py` mentioned in section "**Via `docker-compose`**"

  Follow the steps to deploy voila-cdn-replica server on an existing Kubernetes cluster. The yaml config files are provided under `k8s` folder. Change the content if necessary.

  ``` shell

  # Load settings.py into a config map
  kubectl create configmap voila-cdn-replica --from-file=/path/to/settings.py

  # Create a deployment
  kubectl create -f k8s/voila-cdn-replica.deployment.yml

  # Create a service
  kubectl create -f k8s/voila-cdn-replica.svc.yml

  # Expose the service by executing on master node:
  kubectl port-forward svc/voila-cdn-replica 29000:19000 --address 0.0.0.0

  ```

  Now visit your master node and Voilà!


  ### Manually

  If you prefer running the docker containers manually, then the following command may be what you need. You still need to create the `settings.py` and specify the secret key within.

  ``` shell

    docker run -d \
      --name vcdn-rep-c1 \
      -v /path/2/media:/www/voila-CDN-replica/media \
      -v /path/2/settings.py:/www/voila-CDN-replica/replicaServer/settings.py \
      valorad/voila-cdn-replica

    # inspect id
    docker inspect vcdn-ori-c1 | grep '"IPAddress"' | head -n 1
    
  ```