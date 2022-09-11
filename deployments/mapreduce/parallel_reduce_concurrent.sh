#!/usr/bin/env bash
set -euox pipefail
main() {
  local TIMEOUT="timeout -k 2s 180s"
  ${TIMEOUT} ./comain.exe ./../../assets/mapreduce/pg*.txt &
  sleep 1
  ${TIMEOUT} ./worker.exe ./rtiming.so &
  ${TIMEOUT} ./worker.exe ./rtiming.so
  # refer https://github.com/koalaman/shellcheck/wiki/SC2126
  local NT=$(cat mr-out* | grep -c '^[a-z] 2' | sed 's/ //g')
  if [[ "${NT}" -lt "2" ]]; then
    echo '---' too few parallel reduces.
    echo '---' reduce parallelism test: FAIL
    failed_any=1
  else
    echo '---' reduce parallelism test: PASS
  fi

  wait
}
main
