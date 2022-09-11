SHELL := /bin/bash
RACE=-race
GOC=go
GOOS=GOOS=linux
GO=${GOOS} ${GOC}
TIMEOUT=timeout -k 2s 180s

plugin_crash: ./../../pkg/mapreduce/crash/main/crash.go
	${GOC} build \
	  ${RACE} \
	  -buildmode=plugin \
	  -o ./crash.so \
	  ./../../pkg/mapreduce/crash/main/crash.go

plugin_early_exit: ./../../pkg/mapreduce/early_exit/main/early_exit.go
	${GOC} build \
	  ${RACE} \
	  -buildmode=plugin \
	  -o ./early_exit.so \
	  ./../../pkg/mapreduce/early_exit/main/early_exit.go

plugin_indexer: ./../../pkg/mapreduce/indexer/main/indexer.go
	${GOC} build \
	  ${RACE} \
	  -buildmode=plugin \
	  -o ./indexer.so \
	  ./../../pkg/mapreduce/indexer/main/indexer.go

plugin_jobcount: ./../../pkg/mapreduce/jobcount/main/jobcount.go
	${GOC} build \
	  ${RACE} \
	  -buildmode=plugin \
	  -o ./jobcount.so \
	  ./../../pkg/mapreduce/jobcount/main/jobcount.go
plugin_mtiming: ./../../pkg/mapreduce/mtiming/main/mtiming.go
	${GOC} build \
	  ${RACE} \
	  -buildmode=plugin \
	  -o ./mtiming.so \
	  ./../../pkg/mapreduce/mtiming/main/mtiming.go
plugin_nocrash: ./../../pkg/mapreduce/nocrash/main/nocrash.go
	${GOC} build \
	  ${RACE} \
	  -buildmode=plugin \
	  -o ./nocrash.so \
	  ./../../pkg/mapreduce/nocrash/main/nocrash.go
plugin_rtiming: ./../../pkg/mapreduce/rtiming/main/rtiming.go
	${GOC} build \
	  ${RACE} \
	  -buildmode=plugin \
	  -o ./rtiming.so \
	  ./../../pkg/mapreduce/rtiming/main/rtiming.go
plugin_wc: ./../../pkg/mapreduce/wc/main/wc.go
	${GOC} build \
	  ${RACE} \
	  -buildmode=plugin \
	  -o ./wc.so \
	  ./../../pkg/mapreduce/wc/main/wc.go

sequential: ./../../app/mapreduce/sequence/mrsequential.go
	${GOC} build \
		${RACE} \
		-o sequential.exe \
		./../../app/mapreduce/sequence/mrsequential.go

coordinator:  ./../../app/mapreduce/mrcoordinator.go ./../../pkg/mapreduce/coordinator.go ./../../pkg/mapreduce/rpc.go ./../../pkg/mapreduce/worker.go
	${GOC} build \
		${RACE} \
		-o comain.exe \
		./../../app/mapreduce/mrcoordinator.go

worker:  ./../../app/mapreduce/mrworker.go  ./../../pkg/mapreduce/coordinator.go ./../../pkg/mapreduce/rpc.go ./../../pkg/mapreduce/worker.go
	${GOC} build \
		${RACE} \
		-o worker.exe \
		./../../app/mapreduce/mrworker.go

.PHONY: plugin clean wc-sequen build

plugin: plugin_crash plugin_early_exit plugin_indexer plugin_jobcount plugin_mtiming plugin_nocrash plugin_rtiming plugin_wc

### wc begin
wc-sequen: sequential plugin_wc
	./sequential.exe \
		./wc.so \
		./../../assets/mapreduce/pg*.txt \
	&& sort ./mr-out-0 > ./../../assets/mapreduce/mr-correct-wc.txt \
	&& rm ./mr-out-0

wc-comain: coordinator plugin_wc
	./comain.exe ./../../assets/mapreduce/pg*.txt

wc-worker: worker plugin_wc
	./worker.exe ./wc.so

.ONESHELL:
wc-compare:
	cd ./../../assets/mapreduce/
	if cmp mr-wc-all mr-correct-wc.txt ;then
		echo '---' wc test: PASS
	else
		echo '---' wc output is not the same as mr-correct-wc.txt
		echo '---' wc test: FAIL
	fi
	cd ./../../deployments/mapreduce

.ONESHELL:
wc-concurrent: wc-sequen worker coordinator plugin_wc
	./wccon.sh
### wc end

### index begin

index-sequen: sequential plugin_indexer
	./sequential.exe \
		./indexer.so \
		./../../assets/mapreduce/pg*.txt \
	&& sort ./mr-out-0 > ./../../assets/mapreduce/mr-correct-indexer.txt \
	&& rm ./mr-out-0

index-comain: coordinator plugin_indexer
	./comain.exe ./../../assets/mapreduce/pg*.txt

index-worker: worker plugin_indexer
	./worker.exe ./indexer.so
.ONESHELL:
index-compare:
	cd ./../../assets/mapreduce/
	if cmp mr-indexer-all mr-correct-indexer.txt ;then
		echo '---' indexer test: PASS
	else
		echo '---' indexer output is not the same as mr-correct-wc.txt
		echo '---' indexer test: FAIL
	fi
	cd ./../../deployments/mapreduce

.ONESHELL:
index-concurrent: index-sequen worker coordinator plugin_indexer
	./indexcon.sh
### index end

clean:
	rm ./*.so mr-*-*
	rm /var/tmp/824-mr-*


build: plugin sequential coordinator worker
test: build wc-sequen index-sequen
