#!/bin/bash

GOVERSION=1.5

## Compile with docker
docker run --rm \
  -v $(pwd):/go/src/github.com/bernardolins/clustereasy/ \
  -w go/src/github.com/bernardolins/clustereasy \
  golang:$GOVERSION \
  go get ./... && \
  go build -v
