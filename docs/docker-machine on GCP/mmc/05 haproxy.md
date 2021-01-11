# haproxy


```bash
sudo mkdir -p /etc/haproxy/certs

DOMAIN='cwxstat.io' sudo -E bash -c 'cat /etc/letsencrypt/live/$DOMAIN/fullchain.pem /etc/letsencrypt/live/$DOMAIN/privkey.pem > /etc/haproxy/certs/$DOMAIN.pem'
DOMAIN='dev.cwxstat.io' sudo -E bash -c 'cat /etc/letsencrypt/live/$DOMAIN/fullchain.pem /etc/letsencrypt/live/$DOMAIN/privkey.pem > /etc/haproxy/certs/$DOMAIN.pem'

sudo chmod -R go-rwx /etc/haproxy/certs

```

# Setup

```yaml
# My setup
# Ref: https://cbonte.github.io/haproxy-dconv/1.7/configuration.html
frontend ft_test
  bind *:3000 ssl crt /etc/haproxy/certs/cwxstat.io.pem crt /etc/haproxy/certs/dev.cwxstat.io.pem
  use_backend cwxstat_cert if { ssl_fc_sni cwxstat.io } # content switching based on SNI
  use_backend dev_cwxstat_cert if { ssl_fc_sni dev.cwxstat.io } # content switching based on SNI

backend cwxstat_cert
   balance roundrobin
   server dev 172.18.1.129:80 check


backend dev_cwxstat_cert
   balance roundrobin
   server dev2 172.18.1.128:80/productpage check


```

# Check proxy setup

```bash
haproxy -c -f haproxy.cfg
# start
service haproxy start
service haproxy status

```


nginx
```bash

cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  labels:
    run: nginx
  name: nginx
spec:
  containers:
  - name: nginx 
    image: nginx
    ports:
    - containerPort: 80
EOF

kubectl expose pod nginx --type=LoadBalancer
```





```bash

curl https://dev.cwxstat.io:3000/productpage


kubectl expose pod nginx --type=LoadBalancer

kubectl -n istio-system expose service kiali --type=LoadBalancer

```