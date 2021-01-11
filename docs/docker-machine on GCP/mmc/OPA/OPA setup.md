# OPA

```bash
# OPA cluster

kind delete cluster --name opa 
kind create cluster --name opa --image=quay.io/mchirico/k8s:v1.20.1 --config configs/kind_opa.yaml


```

## Forward ports

```bash

export PORT=$(docker ps --filter "name=opa-control-plane" --format "{{.Ports}}"| sed -e 's/.*://'|sed -e 's/->.*//g')
docker-machine ssh ${KUBE_BUILD_VM} -L ${PORT}:localhost:${PORT} -N &

```


Metal

```bash
# Metal
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.9.5/manifests/namespace.yaml
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.9.5/manifests/metallb.yaml
# On first install only
kubectl create secret generic -n metallb-system memberlist --from-literal=secretkey="$(openssl rand -base64 128)"


kubectl apply -f configs/metal_opa.yaml 


```




## Now installing OPA

```bash

kubectl create -f https://raw.githubusercontent.com/killer-sh/cks-course-environment/master/course-content/opa/gatekeeper.yaml

# Check status

k get crd

# Next: Create  (k get K8sAlwaysDeny) policy

k apply -f https://raw.githubusercontent.com/killer-sh/cks-course-environment/master/course-content/opa/deny-all/all_pod_always_deny.yaml

# Check status

k get K8sAlwaysDeny



```


### Take a look at Github

Take a look at the following github
https://github.com/killer-sh/cks-course-environment/tree/master/course-content/opa/deny-all


Now create it.  Note, this is the template that you can modify.

```bash
k create -f https://raw.githubusercontent.com/killer-sh/cks-course-environment/master/course-content/opa/deny-all/alwaysdeny_template.yaml

# To see status

k get constrainttemplates


```

# Note: Now you can't create a pod.  Try it...

```bash
kubectl run nginx --image=nginx

```


# Get status on this

```bash
k describe k8salwaysdeny

# You'll see 13 violations here.  Why 13?

```



# Change it
```bash
wget https://raw.githubusercontent.com/killer-sh/cks-course-environment/master/course-content/opa/deny-all/alwaysdeny_template.yaml

```

Here's some more detail on the template.

```yaml
apiVersion: templates.gatekeeper.sh/v1beta1
kind: ConstraintTemplate
metadata:
  name: k8salwaysdeny
spec:
  crd:
    spec:
      names:
        kind: K8sAlwaysDeny
      validation:
        # Schema for the `parameters` field
        openAPIV3Schema:
          properties:
            message:
              type: string
  targets:
    - target: admission.k8s.gatekeeper.sh
      rego: |
        package k8salwaysdeny

        violation[{"msg": msg}] {
          1 > 0 # true
          1 > 2 # false... if you have ANY false, we exit.
          msg := input.parameters.message
        }

```