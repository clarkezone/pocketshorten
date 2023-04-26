# End-to-end demo

## Prerequisites

### Environment

In order to follow along you'll need the following:

1. A Kubernetes cluster
2. Default storageclass configured (eg Longhorn)
3. A cloudflare account and domain

### Tools

Ensure the following tools are installed:

1. `helm`
2. `kubectl`
3. `cue`
4. `cloudflared`

## Deploy monitoring

1. (Optional) to install a graphana stack on the cluster, go [here](https://github.com/clarkezone/pocketshorten/tree/main/endtoenddemo/manifests/grafana-stack)

## Deploy test target site to cluster

Ensure that the `cloudflared` cli is logged in by typing `cloudflared tunnel list`

1. Run following commands to prepare and deploy a nginx test website onto the cluster to use as a target of URL shortening:

   ```bash
   # copy manifests for a test website which will be the target of shorten operations
   ./createapplytestsite.sh

   # ensure cloudflare tunnel is created, update the manifests with secrets and tunnel identifiers
   ./createtunnel.sh pocketshortene2edemo-target-tunnel-prod psdemotarget.clarkezone.dev manifests/apply/nginx_simplefile_apply

   # apply prepared manifests to cluster
   kubectl apply -k manifests/apply/nginx_simplefile_apply
   ```

2. Verify that the site is up by visiting [https://psdemotarget.clarkezone.dev](https://psdemotarget.clarkezone.dev)

3. To instpect the manifests as applied to the cluster use `kubectl kustomize manifests/apply/nginx_simplefile_apply`

## Deploy pocketshorten to cluster

Deploy the url shortener application to the cluster. Use the following configuration file to set up a couple of shortcut routes that point to the test website from 1. above.

```json
{
  "values": [
    [
      "tsh",
      "https://psdemotarget.clarkezone.dev/",
      "group",
      "2023-04-22T15:04:05-0700"
    ],
    [
      "lol",
      "https://psdemotarget.clarkezone.dev/canihascheezburger-meme-page.html",
      "group",
      "2023-04-22T15:04:05-0700"
    ],
    [
      "gc",
      "https://psdemotarget.clarkezone.dev/grumpycat-meme-page.html",
      "group",
      "2023-04-22T15:04:05-0700"
    ],
    [
      "gns",
      "https://psdemotarget.clarkezone.dev/gangnamstyle-meme-page.html",
      "group",
      "2023-04-22T15:04:05-0700"
    ],
    ["tm", "https://techmeme.com", "sites", "2023-04-22T15:04:05-0700"],
    ["hn", "https://news.ycombinator.com", "sites", "2023-04-22T15:04:05-0700"]
  ]
}
```

1. To perform deployment, run the following commands:

   ```bash
   # copy pocketshorten manifests to apply directory, copy config files for test deployment
   ./createapplypocketshorten.sh

   # ensure cloudflare tunnel is created for url shortener application, update the manifests with secrets and tunnel identifiers
   ./createtunnel.sh pocketshortene2edemo-tunnel-prod psdemo.clarkezone.dev manifests/apply/pocketshorten_apply/overlay/prod

   # Apply the preparted manifests
   kubectl apply -k manifests/apply/pocketshorten_apply/overlay/prod
   ```

2. Verify shortener is working by visiting [https://psdemo.clarkezone.dev]

3. To instpect the manifests as applied to the cluster use `kubectl kustomize manifests/apply/pocketshorten_apply/overlay/prod`

## Run load

1. Run load on local dev machine.

   ```bash
   k6 run endpoint_prod_variable.js
   ```

2. Import dashboard and view traffic

## Grafana cloud scenario

### Install grafana cloud stack

1. uninstall the local observability stack:

   ```bash
   kubectl delete namespace monitoring
   kubectl delete namespace loki-stack
   ```

2. Create a new Grafana cloud account
3. Connect Data
4. Start sending data
5. scroll down and click agent operator
6. ensure cluster is named `rapi-c4` and namespace is `grafanacloud`
7. open `installgrafanacloud.sh` in VSCODE (to paste from host) and replace `#TODO: insert config from dashboard` with the config
8. IMPORTANT! Update `serviceMonitorSelector` and `podMonitorSelector` with {}
9. Run `installgrafanacloud.sh`

Browse through pages of K8s default dashboards
Look at logs

### Run cloud load using k6 in cloud mode

1. Import pocketshorten dashboard
2. Goto load tab in browser, get login command
3. Run load `k6 cloud endpoint_prod_variable.js`
4. open load test in browser, click on running

### Add an alerting rule

1. From hamburger, go to alerts
2. Click on create new rule
3. click on Code and paste `sum(rate(pocketshorten_frontend_totalops[5m]))`
4. Click preview to see the graph is alerting

Uninstall:

```bash
helm uninstall operator -n grafanacloud
kubectl delete namespace grafanacloud
```

## Azure Kubernetes Service scenario

### Create cluster and retrieve credentials

```bash
# Login to Azure
az login

# ensure that the correct subscription is selected
az account show

# create AKS cluster
az group create -n shortenaks-rg --location westus
az aks create -g shortenaks-rg -n shortencluster --generate-ssh-keys
az aks get-credentials --admin --name shortencluster --resource-group shortenaks-rg

# Get Kube Config
az aks get-credentials --admin --name shortencluster --resource-group shortenaks-rg
```

### Install Grafana Cloud observability

1. open installgrafanacloud.sh in vscode
2. Replace rapi-c4
3. `./installgrafanacloud.sh`

### Install Pocketshorten

```bash
   kubectl apply -k manifests/apply/pocketshorten_apply/overlay/prod
```

### Uninstall

```bash
helm uninstall operator -n grafanacloud
kubectl delete namespace grafanacloud
kubectl delete namespace pocketshorten
```

### Reset

```bash
helm uninstall operator -n grafanacloud
kubectl delete namespace grafanacloud
kubectl delete namespace pocketshorten
# DELETE AKS CLUSTER!

k config use-context rapi-c4
helm uninstall operator -n grafanacloud
kubectl delete namespace grafanacloud
kubectl delete namespace pocketshorten
kubectl delete namespace nginxsimple

az account show

```

Delete stacks from: `https://grafana.com/orgs/jeclarke`
