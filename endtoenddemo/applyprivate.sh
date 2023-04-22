#!/bin/bash
set -e
cue vet pocketshorten.json ../testfiles/schema.cue

rm -rf manifest/apply/layered_viper_apply
rm -rf manifest/apply/nginx_simplefile

mkdir -p manifests/apply/layered_viper_apply
cp -r ../k8s/layered_viper/* manifests/apply/layered_viper_apply/.
cp pocketshorten.json manifests/apply/layered_viper_apply/base/.

mkdir -p manifests/apply/layered_viper_apply/overlay/prod/config
mkdir -p manifests/apply/layered_viper_apply/overlay/prod/secrets

./createtunnel.sh pocketshortene2edemo-tunnel-prod psdemo.clarkezone.dev manifests/apply/layered_viper_apply/overlay/prod

# For the nginx test target deployment
mkdir -p manifests/apply/nginx_simplefile_apply
cp -r manifests/nginx_simplefile/* manifests/apply/nginx_simplefile_apply
./createtunnel.sh pocketshortene2edemo-target-tunnel-prod psdemotarget.clarkezone.dev manifests/apply/nginx_simplefile_apply
