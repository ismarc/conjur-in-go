#!/usr/bin/env bash

go generate ./cmd/conjurctl
go run ./cmd/conjurctl server -m