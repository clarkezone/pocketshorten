# End-to-end demo

## Prerequisites

1. helm
2. kubectl
3. cue
4. cloudflared
5. A kubernetes cluster
6. Default storageclass configured (eg Longhorn)
7. A cloudflare account and domain
8.

## Install

1. Install monitoring stack todo link
2. ./createapply.sh
3. kubectl apply -k manifests/apply/pocketshorten_apply/overlay/prod
4. kubectl apply -k manifests/apply/nginx_simplefile_apply

## Run load

1. k6 run endpoint_prod_variable.js (switch k9s to nodes)

Todo:

1. Add two more pages to nginx site and make load test randomly pick those
2. break out cloudflare tunnel into this readme / separate script
3. Is there a way of showing peek r/s in last hour?
4. tidy up this readme with bash

Grafana cloud scenario

1. walk through grafana cloud steps

Azure Kubernetes Service scenario

1. walk through the steps
