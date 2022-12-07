# Pocketshorten

[![License](https://img.shields.io/github/license/clarkezone/pocketshorten.svg)](https://github.com/clarkezone/pocketshorten/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/clarkezone/pocketshorten)](https://goreportcard.com/report/github.com/clarkezone/pocketshorten)
[![Build and Tests](https://github.com/clarkezone/pocketshorten/workflows/run%20tests/badge.svg)](https://github.com/clarkezone/pocketshorten/actions?query=workflow%3A%22run+tests%22) [![Coverage Status](https://coveralls.io/repos/github/clarkezone/pocketshorten/badge.svg?branch=main)](https://coveralls.io/github/clarkezone/pocketshorten?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/clarkezone/pocketshorten.svg)](https://pkg.go.dev/github.com/clarkezone/pocketshorten)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/6231/badge)](https://bestpractices.coreinfrastructure.org/projects/6231)
[![GitHub release](https://img.shields.io/github/release/clarkezone/pocketshorten.svg?style=flat-square)](https://github.com/clarkezone/pocketshorten/releases)
![Total Downloads](https://img.shields.io/github/downloads/clarkezone/pocketshorten/total?logo=github&logoColor=white)

# project state

Bootstrapping infra:

- [x] Local build
- [x] Local build from dockerfile
- [x] Local precommit / linting / githook
- [x] Fix name to be consistent
- [x] Implement CI build, publish to docker
- [x] Fix CI artifacts (eg code coverage)
- [x] Fix repo badges
- [x] Best practices from mail log
- [x] Metrics exposed on independent port
- [x] k8s public
- [x] Service Monitor
- [ ] Add load tester
- [ ] Dashboard for telemetry
- [ ] Hello world comes from microservice with switch / env for mode in single binary
- [ ] Cleanup names
- [ ] Fix double counting of metrics
- [ ] k8s local override with staging / prod
- [ ] Instructions for prerequs
- [ ] Fork and make template (rename target to be new baseline)
- [ ] Turn on PR enforcement, protect main branch

Template backlog

- Add k8s manifest scanner for best practices (PDB, CPU/MEM requests and limits)
- Add DT

Bring over shortening prototype

# Building

## Install prerequisites

1. golang 1.9x
2. make
3. TODO rest of items
4. `precommit-installhooks`

# Backlog

- Add vscode devcontaienr
- complete openssf best practices
- Add minimal viable covernance

# Breating a new release

```bash
git tag -a v0.0.1 -m "helloinfra"
git push origin v0.0.1
gh release create
```
