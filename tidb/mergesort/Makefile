.PHONY: all

run = 5

all: clean test bench

test:
	go test

bench:
	go test -bench Benchmark -run xx -count $(run) -benchmem

analysis: 
	go test -bench Benchmark -run xx -count $(run) -benchmem -cpuprofile cpu.prof -trace=trace.out

profile:
	go test -bench Benchmark -run xx -count $(run) -benchmem -cpuprofile cpu.prof

trace:
	go test -bench Benchmark -run xx -count $(run) -benchmem -trace=trace.out

clean:
	rm cpu.prof trace.out mergesort.test

