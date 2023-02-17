#!/bin/bash
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
LOADPATH="$SCRIPT_DIR/endpoint_stage_variable.js"
k6 run """$LOADPATH"""
