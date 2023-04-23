#!/bin/bash
set -e
cue vet pocketshorten.json ../testfiles/schema.cue

rm -rf manifest/apply/pocketshorten_apply

mkdir -p manifests/apply/pocketshorten_apply
cp -r ../k8s/layered_viper/* manifests/apply/pocketshorten_apply/.
cp pocketshorten.json manifests/apply/pocketshorten_apply/base/.

mkdir -p manifests/apply/pocketshorten_apply/overlay/prod/tunnelconfig
mkdir -p manifests/apply/pocketshorten_apply/overlay/prod/tunnelsecrets
