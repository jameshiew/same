#!/usr/bin/env sh

set -eoux pipefail

go install -ldflags "-X main.version=dev-$(date +%s)"
