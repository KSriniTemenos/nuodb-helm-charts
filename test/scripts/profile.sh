#!/usr/bin/env bash

# optionally set environment variable for
# 'google', 'azure', or 'amazon'.

: ${PLATFORM:="google"}
: ${CLUSTER_NAME:="helmtest"}
: ${ZONE:="us-central1"}
: ${TARGET_NAMESPACE:="nuodb"}

export PLATFORM
export CLUSTER_NAME
export TARGET_NAMESPACE
