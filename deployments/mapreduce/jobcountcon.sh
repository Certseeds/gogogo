#!/usr/bin/env bash
set -euox pipefail
main() {
  local TIMEOUT="timeout -k 2s 180s"
  ${TIMEOUT} ./comain.exe ./../../assets/mapreduce/pg*.txt &
  local pid=$!
  sleep 1
  $TIMEOUT ./worker.exe ./jobcount.so &
  $TIMEOUT ./worker.exe ./jobcount.so
  $TIMEOUT ./worker.exe ./jobcount.so &
  $TIMEOUT ./worker.exe ./jobcount.so
  wait ${pid}
  NT=$(cat mr-out* | awk '{print $2}')
  if [ "$NT" -eq "8" ]; then
    echo '---' job count test: PASS
  else
    echo '---' map jobs ran incorrect number of times "($NT != 8)"
    echo '---' job count test: FAIL
  fi
}
main
