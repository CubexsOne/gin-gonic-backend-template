#!/bin/sh

ENV=$1

cd /app
go install github.com/cortesi/modd/cmd/modd

go mod download && modd --file=modd-$ENV.conf