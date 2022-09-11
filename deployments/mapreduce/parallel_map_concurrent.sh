#!/usr/bin/env bash
set -euox pipefail
main() {
  local TIMEOUT="timeout -k 2s 180s"
  ${TIMEOUT} ./comain.exe ./../../assets/mapreduce/pg*.txt &
  sleep 1
  ${TIMEOUT} ./worker.exe ./mtiming.so &
  ${TIMEOUT} ./worker.exe ./mtiming.so
  # refer https://github.com/koalaman/shellcheck/wiki/SC2126
  local NT=$(cat mr-out* | grep -c '^times-' | sed 's/ //g')
  if [[ "$NT" != "2" ]]; then
    echo '---' saw "$NT" workers rather than 2
    echo '---' map parallelism test: FAIL
    exit 1
  fi
  if cat mr-out* | grep '^parallel.* 2' >/dev/null; then
    echo '---' map parallelism test: PASS
  else
    echo '---' map workers did not run in parallel
    echo '---' map parallelism test: FAIL
    exit 2
  fi
  wait
}
main
