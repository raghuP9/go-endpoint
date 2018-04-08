#!/bin/bash

TARGET_HOST=${TARGET_HOST:-gitlab.com}
TARGET_PROTO=${TARGET_PROTO:-https}
TARGET_INTERVAL=${TARGET_INTERVAL:-300}
GO_ENVIRONMENT=${GO_ENVIRONMENT:-production}

if [ $GO_ENVIRONMENT == "development" ]
then
  /go/bin/go-endpoint -host $TARGET_HOST -protocol $TARGET_PROTO -interval $TARGET_INTERVAL --insecure --show
else
  /go/bin/go-endpoint -host $TARGET_HOST -protocol $TARGET_PROTO -interval $TARGET_INTERVAL
fi

