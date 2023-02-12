#!/bin/bash
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
LOADPATH="$SCRIPT_DIR/endpoint_local.js"
k6 run """$LOADPATH""" --duration 30s --vus 200
