#!/usr/bin/env bash

# can be used to test wait commands
# you should update the service.json Script location
# to point to your local code, or make this smarter here
curl -v -X PUT -d @service.json 127.0.0.1:8500/v1/agent/service/register
