# Pocketshorten

[![License](https://img.shields.io/github/license/clarkezone/pocketshorten.svg)](https://github.com/clarkezone/pocketshorten/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/clarkezone/pocketshorten)](https://goreportcard.com/report/github.com/clarkezone/pocketshorten)
[![Build and Tests](https://github.com/clarkezone/pocketshorten/workflows/run%20tests/badge.svg)](https://github.com/clarkezone/pocketshorten/actions?query=workflow%3A%22run+tests%22) [![Coverage Status](https://coveralls.io/repos/github/clarkezone/pocketshorten/badge.svg?branch=main)](https://coveralls.io/github/clarkezone/pocketshorten?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/clarkezone/pocketshorten.svg)](https://pkg.go.dev/github.com/clarkezone/pocketshorten)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/6231/badge)](https://bestpractices.coreinfrastructure.org/projects/6231)
[![GitHub release](https://img.shields.io/github/release/clarkezone/pocketshorten.svg?style=flat-square)](https://github.com/clarkezone/pocketshorten/releases)
![Total Downloads](https://img.shields.io/github/downloads/clarkezone/pocketshorten/total?logo=github&logoColor=white)

## Project State

MVP completed as of 0.0.8

Cleanup

- [ ] Update readme with pre-requs and full instructions
- [ ] complete dashboards

## Phase 3

- [ ] Push-based config updates sourced from microservice
  - [ ] logging for microservice calls
- [ ] Pocketbase or FowlerTodo

## Phases

Phase 1 Infra bringup: CI/CD in place
Phase 2 MVP: Pocketshorten runs natively or in container, state from viper config, metrics and k8s manifests with best practices
Phase 3 UI for modifying shorten routes

## Building

### Install prerequisites

1. golang 1.9x
2. make
3. kubectl
4. krew: [https://krew.sigs.k8s.io/docs/user-guide/quickstart/](https://krew.sigs.k8s.io/docs/user-guide/quickstart/)
5. kube-score: [https://github.com/zegl/kube-score](https://github.com/zegl/kube-score)
6. TODO rest of items
7. `precommit-installhooks`

## Running locally

1. Build using instructions above
2. Start local executable
   `./bin/pocketshorten servefrontend --config testfiles/.pocketshorten.json --loglevel=debug`
3. View telemetry TODO
4. Run local load TODO. TODO need a target that isn't exposing me

## Running locally in Docker

Here you will start a nginx test webserver to act as
a target for the redirect operation and then start an instance
of pocketshorten running in docker.

1. Start docker `docker run --rm -d -p 8080:80 --name web nginx`
2. Create a test configuration file that sets up a redirect rule to point to above nginx server

   ```bash
   cat <<EOF > testfiles/redirectTest.json
   {
     "values": [
       ["nginxlocal", "http://0.0.0.0:8080", "testgroup", "2023-01-02T15:04:05-0700"]
     ]
   }
   EOF
   ```

3. Start pocketshorten in docker:

   ```bash
      docker run --rm -d -p 8090:8090 -p 8095:8095 -v ${PWD}/testfiles:/testfiles -e LOGLEVEL=debug --name web nginx registry.hub.docker.com/clarkezone/pocketshorten:main servefrontend --config /testfiles/redirectTest.json
   ```

4. Read telemetry endpoint and look at the lookup statistics

   ```bash
   curl -s localhost:8095 | grep pocketshorten_frontend_total_lookups{
   ```

5. Attempt to resolve the nginxlocal destination:

   ```bash
   curl localhost:8090/nginxlocal
   <a href="http://0.0.0.0:8080">Moved Permanently</a>.

   ```

   You should see a http redirect instructions. If you tell curl to follow redirects you should see the output of the running nginx server returned

   ```bash
   curl -L loclhost:8090/nginxlocal

   <!DOCTYPE html>
   <html>
   <head>
   <title>Welcome to nginx!</title>
   <style>
   html { color-scheme: light dark; }
   ...
   ```

   Repeating the telemetry query should now show the number of lookups at 2

   ```bash
   curl -s localhost:8095 | grep pocketshorten_frontend_total_lookups{
   ```

6. delete containers

   ```bash
   docker stop $(docker ps -q --filter name=nginx )
   docker stop $(docker ps -q --filter name=pocketshorten )
   ```

## Running in Kubernetes with Cloudflare tunnel

pre-requs:

1. k8s cluster (eg k3s)
2. Storage provider (eg longhorn)
3. Prom operator (prometheus, grafana, loki)
4. Cloudflare account
5. Cloudflare API key

steps

1. Cloudflare prep. Configure DNS in cloudflare, create tunnel, put values into config file
2. TODO get other todos from craft

## Backlog

- Add vscode devcontaienr
- complete openssf best practices
- Add minimal viable covernance
- Add container security to pocketshorten
- Add container security to tunnel
- Add limits
- Distributed Throttling / rate limiting

# Creating a new release

```bash
git tag -a v0.0.1 -m "helloinfra"
git push origin v0.0.1
gh release create
```
