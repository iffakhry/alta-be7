# Getting started with docker

Common commands

```bash
# pull image from container registry
docker pull <image-name>

# login to container registry
docker login -u <user-name>

# build image
docker build -t <image-name>[:tag] .

# push image to container registry
docker push <image-name>

# create and run container
docker run -p <host-port>:<container-port> -e <env-name>=<env-value> -v <host-volume>:<container-volume> --name <container-name> <image-name>

# run existing container
docker start <container-name>
docker container start <container-name>

# stop running container
docker stop <container-name>
docker container stop <container-name>

# remove container
docker rm <container-name>
docker container rm <container-name>

# show images
docker image list

# show container
docker container list
docker ps
docker ps -a

# access/run command in a container
docker exec -it <container-name> <command>
docker exec -it mysql bash
```

# Using docker

If you are using docker-desktop, the containers can access host os by using `host.docker.internal` name.

Otherwise, you can use default host IP address: `172.17.0.1`


if you can't run docker on your linux. please run
```
sudo usermod -a -G docker ubuntu
# or: sudo chmod 777 /var/run/docker.sock
# or: sudo su
```