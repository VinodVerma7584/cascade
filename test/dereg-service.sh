#!/usr/bin/env bash

# can be used to test wait commands

curl -v -X PUT 127.0.0.1:8500/v1/agent/service/deregister/redis1
