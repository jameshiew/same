#!/usr/bin/env sh

go build -ldflags "-X main.version=dev-$(date +%s)"
