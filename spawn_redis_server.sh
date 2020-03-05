#!/bin/sh
tmpFile=$(mktemp)
go build -o $tmpFile app/server.go
exec $tmpFile
