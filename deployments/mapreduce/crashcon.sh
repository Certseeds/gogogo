#!/usr/bin/env bash
set -uox pipefail
# pay attention, crash.so will execute os.Exit(1)
# so set -e must close to escape the message
main() {
  local TIMEOUT="timeout -k 2s 180s "
  echo '***' Starting crash test.
  (
    ${TIMEOUT} ./comain.exe ./../../assets/mapreduce/pg*.txt
    touch ./mr-done
  ) &
  # give the coordinator time to create the sockets.
  sleep 1

  # start multiple workers.
  ${TIMEOUT} ./worker.exe ./crash.so  &
  local SOCKNAME
  SOCKNAME=/var/tmp/824-mr-$(id -u)
  (while [[ -e ${SOCKNAME} ]] && [[ ! -f ./mr-done ]]; do
    ${TIMEOUT} ./worker.exe ./crash.so
    sleep 1
  done) &

  (while [[ -e ${SOCKNAME} ]] && [[ ! -f ./mr-done ]]; do
    ${TIMEOUT} ./worker.exe ./crash.so
    sleep 1
  done) &
  while [[ -e ${SOCKNAME} ]] && [[ ! -f ./mr-done ]]; do
    ${TIMEOUT} ./worker.exe ./crash.so
    sleep 1
  done
  wait
  rm ${SOCKNAME}
  sort mr-out* | grep . | tee ./../../assets/mapreduce/mr-crash-all
  if cmp ./../../assets/mapreduce/mr-correct-crash.txt ./../../assets/mapreduce/mr-crash-all ; then
    echo '---' crash test: PASS
  else
    echo '---' crash output is not the same as mr-correct-crash.txt
    echo '---' crash test: FAIL
    failed_any=1
  fi
}
main
