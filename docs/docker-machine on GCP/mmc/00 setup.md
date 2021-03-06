## Docker Macine on GCP

# Reference
https://docs.docker.com/machine/drivers/gce/
and
https://github.com/kubernetes/kubernetes/blob/8725c3bf12cfd3697464136201216fa05dc662d2/build/README.md


# Start Steps (Already Created)
```bash
# export GOOGLE_APPLICATION_CREDENTIALS=/Users/rommel/.googleAuthPath/k8s-dev-cwxstat.json
export GOOGLE_APPLICATION_CREDENTIALS=/Users/rommel/.googleAuthPath/mchirico-k8s.json
export KUBE_BUILD_VM=mchirico-build
export KUBE_BUILD_GCE_PROJECT=mchirico
# Start computer
docker-machine start ${KUBE_BUILD_VM}

# Set environment
eval $(docker-machine env ${KUBE_BUILD_VM})
docker-machine ssh ${KUBE_BUILD_VM} -L 3050:172.18.1.128:80 -N &
export PORT=$(docker ps --filter "name=kind-control-plane" --format "{{.Ports}}"| sed -e 's/.*://'|sed -e 's/->.*//g')
docker-machine ssh ${KUBE_BUILD_VM} -L ${PORT}:localhost:${PORT} -N &

```



Note... `rm -rf ~/.docker/machine/machines/k8s-build`  is needed.

```bash
export GOOGLE_APPLICATION_CREDENTIALS=/Users/rommel/.googleAuthPath/mchirico-k8s.json
export KUBE_BUILD_VM=mchirico-k8s
export KUBE_BUILD_GCE_PROJECT=mchirico

# Note below nonstandard subnet: default2

docker-machine create --driver google \
  --google-project ${KUBE_BUILD_GCE_PROJECT} \
  --google-zone us-central1-f \
  --google-machine-type=n1-standard-4 \
  --google-network=default2 \
  --google-subnetwork=default2 \
  --google-disk-size=100 \
  --google-preemptible=false \
  --google-address=35.202.88.179  \
  --google-disk-type=pd-ssd \
  ${KUBE_BUILD_VM}

```

Next run


```bash
# Set up local docker to talk to that machine
eval $(docker-machine env ${KUBE_BUILD_VM})
```


Build cluster

```bash
kind delete cluster
kind create cluster --image=quay.io/mchirico/k8s:v1.20.1 --config configs/kind_basic.yaml


# OPA cluster

kind delete cluster --name opa 
kind create cluster --name opa --image=quay.io/mchirico/k8s:v1.20.1 --config configs/kind_opa.yaml

```




Find out what port is being used.  Below, you can see it is **59701**

```bash
docker ps -a
CONTAINER ID   IMAGE                          COMMAND                  CREATED         STATUS         PORTS                       NAMES
6f603ae99714   quay.io/mchirico/k8s:v1.20.1   "/usr/local/bin/entr…"   7 minutes ago   Up 7 minutes                               kind-worker
190878786bfa   quay.io/mchirico/k8s:v1.20.1   "/usr/local/bin/entr…"   7 minutes ago   Up 7 minutes   127.0.0.1:59701->6443/tcp   kind-control-plane

```

```bash


# forward local 59701 to that machine so that rsync works
docker-machine ssh ${KUBE_BUILD_VM} -L 59701:localhost:59701 -N &

# You may need to use 172.18.1.129 on 
docker-machine ssh ${KUBE_BUILD_VM} -L 3050:172.18.1.129:80 -N &

export PORT=$(docker ps --filter "name=kind-control-plane" --format "{{.Ports}}"| sed -e 's/.*://'|sed -e 's/->.*//g')
docker-machine ssh ${KUBE_BUILD_VM} -L ${PORT}:localhost:${PORT} -N &

```


## Stop

```bash
export GOOGLE_APPLICATION_CREDENTIALS=/Users/rommel/.googleAuthPath/mchirico-k8s.json
export KUBE_BUILD_VM=mchirico-k8s
export KUBE_BUILD_GCE_PROJECT=mchirico

docker-machine stop ${KUBE_BUILD_VM}


# to remove
docker-machine rm ${KUBE_BUILD_VM}


# to start
docker-machine start ${KUBE_BUILD_VM}

```