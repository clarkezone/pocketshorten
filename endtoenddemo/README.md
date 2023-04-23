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

Ensure that the `cloudflared` cli is logged in.

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

## Deploy pocketshorten to cluster

Deploy the url shortener application to the cluster. Use the following configuration file to set up a couple of shortcut routes that point to the test website from 1. above.

```json
{
  "values": [
    [
      "tsh", # short link
      "https://psdemotarget.clarkezone.dev/", # target url
      "group",
      "2023-04-22T15:04:05-0700"
    ],
    [
      "lol", # short link
      "https://psdemotarget.clarkezone.dev/meme1.html", # target url
      "group",
      "2023-04-22T15:04:05-0700"
    ],
    [
      "m2",
      "https://psdemotarget.clarkezone.dev/meme2.html",
      "group",
      "2023-04-22T15:04:05-0700"
    ],
    [
      "m3",
      "https://psdemotarget.clarkezone.dev/meme3.html",
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

## Run load

1. Run load on local dev machine.

   ```bash
   k6 run endpoint_prod_variable.js (switch k9s to nodes)
   ```

## Grafana cloud scenario

1. walk through grafana cloud steps

## Azure Kubernetes Service scenario

1. walk through the steps
