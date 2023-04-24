#!/bin/bash
set -e

rm -rf manifest/apply/nginx_simplefile_apply

# For the nginx test target deployment
mkdir -p manifests/apply/nginx_simplefile_apply
cp -r manifests/nginx_simplefile/* manifests/apply/nginx_simplefile_apply
