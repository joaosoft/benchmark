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
* https://github.com/olivere/elastic by olivere

```   
/* with joaosoft */
goos: darwin
goarch: amd64
pkg: benchmark/elastic/joaosoft
BenchmarkJoaosoftElastic-4   	       1	45473183048 ns/op
PASS
ok  	benchmark/elastic/joaosoft	45.497s


/* with olivere */
goos: darwin
goarch: amd64
pkg: benchmark/elastic/olivere
BenchmarkGocraftElastic-4   	       1	49515634802 ns/op
PASS
ok  	benchmark/elastic/olivere	49.544s
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
