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

1. k6 run endpoint_prod_variable.js

Todo:

1. fix k6 success
2. Add two more pages to nginx site and make load test randomly pick those
3. Is there a way of showing peek r/s in last hour?
4. tidy up this readme with bash
5. talk through grafana cloud steps
