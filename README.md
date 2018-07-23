# benchmark
[![Build Status](https://travis-ci.org/joaosoft/benchmark.svg?branch=master)](https://travis-ci.org/joaosoft/benchmark) | [![codecov](https://codecov.io/gh/joaosoft/benchmark/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/benchmark) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/benchmark)](https://goreportcard.com/report/github.com/joaosoft/benchmark) | [![GoDoc](https://godoc.org/github.com/joaosoft/benchmark?status.svg)](https://godoc.org/github.com/joaosoft/benchmark/app)

A benchmark to compare my personal project with others.

## Benchmark
At the moment i have the following benchmarks
> Logger
* https://github.com/joaosoft/logger by joaosoft
* https://gitlab.com/vredens/go-logger by vredens

```   
/* with joaosoft */
goos: darwin
goarch: amd64
pkg: benchmark/logger/joaosoft
BenchmarkJoaosoftLoggerWithFormatedMessage-4


/* with vredens */
goos: darwin
goarch: amd64
pkg: benchmark/logger/vredens
BenchmarkVredensLoggerWithFormatedMessage-4
```

> Elastic
* https://github.com/joaosoft/elastic by joaosoft
* https://gitlab.com/gocraft by gocraft

```   
/* with joaosoft */
goos: darwin
goarch: amd64
pkg: benchmark/elastic/joaosoft
BenchmarkJoaosoftElastic-4   	       1	2822202295 ns/op
PASS
ok  	benchmark/elastic/joaosoft	2.839s


/* with gocraft */
goos: darwin
goarch: amd64
pkg: benchmark/elastic/gocraft
BenchmarkGocraftElastic-4   	       1	2989449460 ns/op
PASS
ok  	benchmark/elastic/gocraft	3.014s
```

## Running
on each package ...
```
go test -bench .

```

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
