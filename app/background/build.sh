#!/usr/bin/env bash
set -euox pipefail
main() {
  if [[ ! -f "go.mod" ]]; then
    echo "should invoke in base path"
    exit 1
  fi
  DOCKER_BUILDKIT=1 docker build \
    --progress plain \
    -f $(pwd)/app/background/build.Dockerfile \
    -t gogogo/background:app \
    --target PRODUCT \
    .
  DOCKER_BUILDKIT=1 docker build \
    -f $(pwd)/app/background/build.Dockerfile \
    --output $(pwd)/app/background/ \
    --target EXPORT \
    .
}
main
