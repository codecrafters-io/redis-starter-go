#!/bin/sh

echo "Spawning Redis server..."
exec go run app/server.go