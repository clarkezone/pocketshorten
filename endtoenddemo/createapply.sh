#!/bin/bash
set -e
cue vet pocketshorten.json ../testfiles/schema.cue

rm -rf manifest/apply

mkdir -p manifests/apply/pocketshorten_apply
cp -r ../k8s/layered_viper/* manifests/apply/pocketshorten_apply/.
cp pocketshorten.json manifests/apply/pocketshorten_apply/base/.

mkdir -p manifests/apply/pocketshorten_apply/overlay/prod/tunnelconfig
mkdir -p manifests/apply/pocketshorten_apply/overlay/prod/tunnelsecrets

# For the nginx test target deployment
mkdir -p manifests/apply/nginx_simplefile_apply
cp -r manifests/nginx_simplefile/* manifests/apply/nginx_simplefile_apply
