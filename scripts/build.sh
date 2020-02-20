#!/usr/bin/env sh

set -eoux pipefail

go build -ldflags "-X main.version=dev-$(date +%s)"
