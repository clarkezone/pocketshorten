#!/bin/bash
curl -LO https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/v0.52.0/bundle.yaml
sed -i 's/namespace: default/namespace: monitoring/g' bundle.yaml
