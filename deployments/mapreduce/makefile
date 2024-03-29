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
	chmod +x ./wccon.sh
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
	chmod +x ./indexcon.sh
	./indexcon.sh
### index end

### mapper parallel begin
parallel-map-comain: coordinator plugin_mtiming
	./comain.exe ./../../assets/mapreduce/pg*.txt

parallel-map-worker: worker plugin_mtiming
	./worker.exe ./mtiming.so

.ONESHELL:
parallel-map-concurrent: coordinator worker plugin_mtiming
	chmod +x ./parallel_map_concurrent.sh
	./parallel_map_concurrent.sh
### mapper paraller end

### reducer parallel begin
parallel-reduce-comain: coordinator plugin_rtiming
	./comain.exe ./../../assets/mapreduce/pg*.txt

parallel-reduce-worker: worker plugin_rtiming
	./worker.exe ./rtiming.so

.ONESHELL:
parallel-reduce-concurrent: coordinator worker plugin_rtiming
	chmod +x ./parallel_reduce_concurrent.sh
	./parallel_reduce_concurrent.sh
### reducer paraller end

### jobcount begin
jobcount-comain: coordinator plugin_jobcount
	./comain.exe ./../../assets/mapreduce/pg*.txt

jobcount-worker: worker plugin_jobcount
	./worker.exe ./jobcount.so

.ONESHELL:
jobcount-concurrent: coordinator worker plugin_jobcount
	chmod +x ./jobcountcon.sh
	./jobcountcon.sh
### jobcount end

### early_exit begin
earlyexit-comain: coordinator plugin_early_exit
	./comain.exe ./../../assets/mapreduce/pg*.txt

earlyexit-worker: worker plugin_early_exit
	./worker.exe ./early_exit.so

.ONESHELL:
earlyexit-concurrent: coordinator worker plugin_early_exit
	chmod +x ./early_exit_concurrent.sh
	./early_exit_concurrent.sh
### early_exit end

### crash begin
crash-comain: coordinator plugin_crash
	./comain.exe ./../../assets/mapreduce/pg*.txt

crash-worker: worker plugin_crash
	./worker.exe ./early_exit.so

.ONESHELL:
crash-sequen: sequential plugin_nocrash
	./sequential.exe \
		./nocrash.so \
		./../../assets/mapreduce/pg*.txt \
	&& sort ./mr-out-0 > ./../../assets/mapreduce/mr-correct-crash.txt \
	&& rm ./mr-out-0

.ONESHELL:
crash-concurrent: coordinator worker plugin_crash crash-sequen
	chmod +x ./crashcon.sh
	./crashcon.sh
### crash end


clean:
	rm ./*.so
	rm ./*.exe
	rm ./mr-*-*
	rm /var/tmp/824-mr-*
	rm ./mr-worker-map-*
	rm ./mr-worker-reduce-*
	rm ./mr-worker-jobcount-*
	rm ./anydone*
	rm ./mr-done



build: plugin sequential coordinator worker
pretest: build wc-sequen index-sequen crash-sequen

.ONESHELL:
test: pretest
	make clean || true
	make wc-concurrent
	make clean || true
	make index-concurrent
	make clean || true
	make parallel-map-concurrent
	make clean || true
	make parallel-reduce-concurrent
	make clean || true
	make jobcount-concurrent
	make clean || true
	make earlyexit-concurrent
	make clean || true
	make crash-concurrent
