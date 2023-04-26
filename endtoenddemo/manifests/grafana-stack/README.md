# Install monitoring stack

## Install loki and promtail

```bash
kubectl create namespace loki-stack
helm upgrade --install loki --namespace=loki-stack grafana/loki-stack
```

## Install prometheus operator

```bash
# Fetch bundle manifest and update namespace
curl -LO https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/v0.52.0/bundle.yaml
sed -i 's/namespace: default/namespace: monitoring/g' bundle.yaml

# Apply bundle into monitoring namespace
kubectl create namespace monitoring
kubectl apply -f bundle.yaml --force-conflicts=true --server-side
```

## Apply minimal Prometheus / Alertmanager / Grapha stack using prometheus operator

```bash
kubectl apply -k manifests/grafana-stack/monitoring/overlay/production
```

# Visit Grafana page

```bash
kubectl port-forward -n monitoring services/grafana-service 3000:3000 --address 0.0.0.0
```

Point browser at `http://localhost:3000`

# Remove monitoring stack

## Remove loki

```bash
# loki
helm uninstall -n loki-stack loki
kubectl delete namespace loki-stack

# monitoring
kubectl delete namespace monitoring
```
