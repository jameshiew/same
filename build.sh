#!/usr/bin/env sh
go build -v -ldflags "-X main.version=dev-$(date +%s)" -o same main.go
