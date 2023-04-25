#!/bin/bash
set -e
NAMESPACE=grafanacloud

kubectl create namespace $NAMESPACE

helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
helm install operator --create-namespace grafana/grafana-agent-operator -n $NAMESPACE

cat <<'EOF' | NAMESPACE=grafanacloud /bin/sh -c 'kubectl apply -n $NAMESPACE -f -'

EOF
