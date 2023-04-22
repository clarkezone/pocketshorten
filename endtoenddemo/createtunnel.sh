#!/bin/bash
set -e
export tunnelname="pocketshortene2edemo-tunnel-prod"
export domain="psdemo.clarkezone.dev"
export k8sroot=""
tunnelid=$(cloudflared tunnel list | sed -n "/^.*${tunnelname}[[:space:]].*$/p" | awk '{print $1}')
if [ -z "$tunnelid" ];
then
  echo "Tunnel not found, creating"
  cloudflared tunnel create $tunnelname
  cloudflared tunnel route dns $tunnelname $domain
fi
tunnelid=$(cloudflared tunnel list | sed -n "/^.*${tunnelname}[[:space:]].*$/p" | awk '{print $1}')
echo "$tunnelid"
cp "$HOME/.cloudflared/$tunnelid.json" manifests/layered_viper_apply/overlay/prod/secrets/.
cp "$HOME/.cloudflared/cert.pem" manifests/layered_viper_apply/overlay/prod/secrets/.
sed -i 's/TUNNEL_NAME/'"$tunnelname"'/; s/TUNNEL_ID/'"$tunnelid"'/; s/HOSTNAME/'"$domain"'/' manifests/layered_viper_apply/overlay/prod/config/config.yaml
