# End-to-end demo

## Prerequisites

1. helm
2. kubectl
3. cue
4. cloudflared
5. A kubernetes cluster
6. Default storageclass configured (eg Longhorn)
7. A cloudflare account and domain

## Install

1. Install monitoring stack todo link
2. Run `createapply.sh` script to copy and update manifests ready for deployment, ensure tunnels created for `nginx` and `pocketshorten`

   ```bash
   ./createapply.sh
   ./createtunnel.sh pocketshortene2edemo-tunnel-prod psdemo.clarkezone.dev manifests/apply/pocketshorten_apply/overlay/prod
   ./createtunnel.sh pocketshortene2edemo-target-tunnel-prod psdemotarget.clarkezone.dev manifests/apply/nginx_simplefile_apply
   ```

3. Apply the prepared manifests:

   ```bash
   kubectl apply -k manifests/apply/pocketshorten_apply/overlay/prod
   kubectl apply -k manifests/apply/nginx_simplefile_apply
   ```

## Run load

1. k6 run endpoint_prod_variable.js (switch k9s to nodes)

Todo:

1. Add two more pages to nginx site and [DONE]
2. make load test randomly pick those
3. break out cloudflare tunnel into this readme / separate script [DONE]
4. Is there a way of showing peek r/s in last hour?
5. tidy up this readme with bash

Grafana cloud scenario

1. walk through grafana cloud steps

Azure Kubernetes Service scenario

1. walk through the steps
