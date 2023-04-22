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

1. Install monitoring stack
2. ./createapply.sh
3. kubectl apply -k manifests/apply/pocketshorten_apply/overlay/prod
