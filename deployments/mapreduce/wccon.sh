#!/usr/bin/env bash
set -euox pipefail
main() {
  local TIMEOUT="timeout -k 2s 180s"
  ${TIMEOUT} ./comain.exe ./../../assets/mapreduce/pg*.txt &
  local pid=$!
  sleep 1
  ${TIMEOUT} ./worker.exe ./wc.so &
  ${TIMEOUT} ./worker.exe ./wc.so &
  ${TIMEOUT} ./worker.exe ./wc.so &
  wait ${pid}
  sort mr-out* | grep . >./../../assets/mapreduce/mr-wc-all
  make wc-compare
}
main
