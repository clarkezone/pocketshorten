#!/bin/bash
k6 run endpoint_stage.js --threshold "http_req_duration{p50,p95,p99}<200"
