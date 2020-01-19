#!/bin/sh
# If no argument only publish the latest
# If release argument is passed it publishes the tagged
# version in the registry. This is meant to be used
# only once per release
set -e
VERSION=$(cat VERSION.TXT)
PUBLISH=false
if [ "$1" == "release" ]; then
  PUBLISH=true
fi

echo "Building/Publishing Docker image go_qmk:latest"
docker build -t docker.pkg.github.com/zekth/go_qmk/go_qmk:latest .
docker push docker.pkg.github.com/zekth/go_qmk/go_qmk:latest

if $PUBLISH; then
  echo "Building/Publishing Docker image go_qmk:$VERSION"
  docker build -t docker.pkg.github.com/zekth/go_qmk/go_qmk:$VERSION .
  docker push docker.pkg.github.com/zekth/go_qmk/go_qmk:$VERSION
fi
