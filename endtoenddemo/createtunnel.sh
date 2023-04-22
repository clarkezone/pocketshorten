#!/bin/bash
set -e
tunnelname="$1"
domain="$2"
target="$3"

tunnelid=$(cloudflared tunnel list | sed -n "/^.*${tunnelname}[[:space:]].*$/p" | awk '{print $1}')
if [ -z "$tunnelid" ];
then
  echo "Tunnel not found, creating"
  cloudflared tunnel create "$tunnelname"
  cloudflared tunnel route dns "$tunnelname" "$domain"
fi
tunnelid=$(cloudflared tunnel list | sed -n "/^.*${tunnelname}[[:space:]].*$/p" | awk '{print $1}')
echo "$tunnelid"
cp "$HOME/.cloudflared/$tunnelid.json" "$target/secrets/."
cp "$HOME/.cloudflared/cert.pem" "$target/secrets/."
sed -i 's/TUNNEL_NAME/'"$tunnelname"'/; s/TUNNEL_ID/'"$tunnelid"'/; s/HOSTNAME/'"$domain"'/' "$target/config/config.yaml"
sed -i 's/TUNNEL_ID/'"$tunnelid"'/' "$target/kustomize.yaml"
