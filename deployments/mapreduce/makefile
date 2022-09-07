SHELL := /bin/bash
RACE=-race
GOC=go
GOOS=GOOS=linux
GO=${GOOS} ${GOC}

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

coordinator:  ./../../app/mapreduce/mrcoordinator.go
	${GOC} build \
		${RACE} \
		-o comain.exe \
		./../../app/mapreduce/mrcoordinator.go

worker:  ./../../app/mapreduce/mrworker.go
	${GOC} build \
		${RACE} \
		-o worker.exe \
		./../../app/mapreduce/mrworker.go

.PHONY: plugin clean sequen-wc build

plugin: plugin_crash plugin_early_exit plugin_indexer plugin_jobcount plugin_mtiming plugin_nocrash plugin_rtiming plugin_wc

sequen-wc: sequential plugin_wc
	./sequential.exe \
		./wc.so \
		./../../assets/mapreduce/pg*.txt \
	&& sort mr-out-0 > ./../../assets/mapreduce/mr-correct-wc.txt \
	&& rm ./mr-out-0

clean:
	rm ./*.so

build: plugin sequential coordinator worker
test: build sequen-wc
