
## User Setup (K8s)
## RBAC

### K8s

For this we'll work in kind-opa

It's section 11: RBAC hardening

```bash
k create ns red
k create ns blue

k -n red create role secret-manager --verb=get --resource=secrets
k -n red create rolebinding secret-manager --role=secret-manager --user=mike

k -n dev-network create role secret-manager --verb=get --verb=list --verb=create --verb=delete --resource=secrets
k -n dev-network create role secret-manager --verb=get --verb=list --verb=create --verb=delete --resource=pods
k -n dev-network create rolebinding secret-manager --role=secret-manager --user=mike

k -n dev-network create role ns-manager  --verb=list  --resource=namespaces
k -n dev-network create rolebinding ns-manager --role=ns-manager --user=mike

# Status
k get rolebinding
k get role




# check permissions
k -n red auth can-i -h
k -n dev-network auth can-i create pods --as mike # yes
k -n red auth can-i get secrets --as jane # yes
k -n red auth can-i list secrets --as jane # no

k -n blue auth can-i list secrets --as jane # yes
k -n blue auth can-i get secrets --as jane # yes

k -n default auth can-i get secrets --as jane #no

# Global
k auth can-i --list


```


# key

```bash
openssl genrsa -out mike.key 2048
openssl req -new -key mike.key -out mike.csr

```

Ref:
https://kubernetes.io/docs/reference/access-authn-authz/certificate-signing-requests/



# Creating base 64
```bash

cat mike.csr |base64
# or linux
cat mike.csr |base64 -w 0
```

Need to paste base64 key in here

```bash
cat <<EOF | kubectl apply -f -
apiVersion: certificates.k8s.io/v1
kind: CertificateSigningRequest
metadata:
  name: mike
spec:
  groups:
  - system:authenticated
  request: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURSBSRVFVRVNULS0tLS0KTUlJQ21UQ0NBWUVDQVFBd1ZERUxNQWtHQTFVRUJoTUNRVlV4RXpBUkJnTlZCQWdNQ2xOdmJXVXRXRMZ24zU0lYYjE1bGdUbnVYQm4vVHhoZjkKLS0tLS1FTkQgQ0VSVElGSUNBVEUgUkVRVUVTVC0tLS0tCg==
  signerName: kubernetes.io/kube-apiserver-client
  usages:
  - client auth
EOF
```


# Now approve

```bash
k get csr
#
k certificate approve mike

#
k get csr mike -o yaml  # Need to base64 decode to mikeout.crt
```


```bash

k config set-credentials mike --client-key=mike.key --client-certificate=mikeout.crt

# To see where you're at
k config get-contexts




k config set-context mike --cluster=kind-opa --user=mike

```

Example output

```bash
k config get-contexts
URRENT   NAME        CLUSTER     AUTHINFO    NAMESPACE
          kind-kind   kind-kind   kind-kind   default
*         kind-opa    kind-opa    kind-opa    default

```

Now set

```bash
k config set-context mike --cluster=kind-opa --user=mike

URRENT   NAME        CLUSTER     AUTHINFO    NAMESPACE
          kind-kind   kind-kind   kind-kind   default
*         kind-opa    kind-opa    kind-opa    default
          mike        kind-opa    mike   

```


```bash

k get clusterRoles -A

```




Now set

```bash
k config use-context mike
# to go back
k config use-context kind-opa

```

k run nginx --image=nginx -n=dev-network



Okay, so what is you want to see everything?

```bash
# This is too much
kubectl create clusterrolebinding mike --clusterrole=edit --user=mike

# This probably won't work.
kubectl create clusterrolebinding mike --clusterrole=list --clusterrole=watch --clusterrole=get --user=mike


```

# Fix

```yaml
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: ns-reader
rules:
- apiGroups: ["","extensions", "apps"] # "" indicates the core API group
  resources: ["*"]
  verbs: ["get", "watch", "list"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: ns-reader
subjects:
- kind: User
  name: mike
roleRef:
  kind: ClusterRole
  name: ns-reader
  apiGroup: rbac.authorization.k8s.io
```  





----------------------------------------------------------

Try the following:

```bash
kubectl apply -f https://k8s.io/examples/pods/security/security-context.yaml

```




