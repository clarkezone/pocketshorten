#!/bin/bash
k6 run endpoint_stage.js --duration 30s --vus 200
