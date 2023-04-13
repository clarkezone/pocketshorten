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
- [ ] Add container security to pocketshorten
- [ ] Add container security to tunnel
- [ ] Add limits
- [ ] complete dashboards

## Phase 3

- [ ] Push-based config updates sourced from microservice
  - [ ] logging for microservice calls
- [ ] Pocketbase or FowlerTodo

## Phases

Phase 1 Infra bringup: CI/CD in place
Phase 2 MVP: Pocketshorten runs natively or in container, state from viper config, metrics and k8s manifests with best practices
Phase 3 UI for modifying shorten routes

## Backlog

- Distributed Throttling / rate limiting

---

# Building

## Install prerequisites

1. golang 1.9x
2. make
3. kubectl
4. krew: [https://krew.sigs.k8s.io/docs/user-guide/quickstart/](https://krew.sigs.k8s.io/docs/user-guide/quickstart/)
5. kube-score: [https://github.com/zegl/kube-score](https://github.com/zegl/kube-score)
6. TODO rest of items
7. `precommit-installhooks`

# Backlog

- Add vscode devcontaienr
- complete openssf best practices
- Add minimal viable covernance

# Creating a new release

```bash
git tag -a v0.0.1 -m "helloinfra"
git push origin v0.0.1
gh release create
```

Starting a url shortener locally:
`./bin/pocketshorten servefrontend --config testfiles/.pocketshorten.json --loglevel=debug`
