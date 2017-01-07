#!/usr/bin/env bash

godep go test -v $(go list ./lang | grep -v /vendor/);
godep go test -v $(go list ./spamc | grep -v /vendor/);
godep go test -v $(go list ./server | grep -v /vendor/);
