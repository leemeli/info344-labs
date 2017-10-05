#!/usr/bin/env bash
# Tell command line: chmod +x build.sh
set -e # If there is an error inany of these commands, stop everything
GOOS=linux go build
docker build -t leemeli/docker-exercise-4 .
go clean

#docker run -d -p 4000:4000 -e PORT=4000 -e FILEPATH=/message/hooray.txt -v /Users/iguest/go/src/github.com/leemeli/info344-labs/docker/docker-exercise-4/message:/message leemeli/docker-exercise-4
