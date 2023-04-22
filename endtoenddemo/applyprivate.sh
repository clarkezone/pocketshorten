#!/bin/bash
set -e
cue vet pocketshorten.json ../testfiles/schema.cue

rm -rf manifest/layered_viper_apply
mkdir -p manifests/layered_viper_apply
cp -r ../k8s/layered_viper/* manifests/layered_viper_apply/.
cp pocketshorten.json manifests/layered_viper_apply/base/.

mkdir -p manifests/layered_viper_apply/overlay/prod/config
mkdir -p manifests/layered_viper_apply/overlay/prod/secrets

./createtunnel.sh pocketshortene2edemo-tunnel-prod psdemo.clarkezone.dev manifests/layered_viper_apply/overlay/prod
