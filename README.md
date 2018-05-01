# go-benchmark
[![Build Status](https://travis-ci.org/joaosoft/go-benchmark.svg?branch=master)](https://travis-ci.org/joaosoft/go-benchmark) | [![codecov](https://codecov.io/gh/joaosoft/go-benchmark/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/go-benchmark) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/go-benchmark)](https://goreportcard.com/report/github.com/joaosoft/go-benchmark) | [![GoDoc](https://godoc.org/github.com/joaosoft/go-benchmark?status.svg)](https://godoc.org/github.com/joaosoft/go-benchmark/app)

A benchmark to compare my personal project with others.

## Benchmark
At the moment i have the following benchmarks
>Logger
* https://github.com/joaosoft/go-log
* https://gitlab.com/vredens/go-logger

```   
/* with joaosoft logger */
   
   // logging to json with one argument
   50000	     29432 ns/op
   
   // logging to json without arguments
   50000	     25919 ns/op


/* with vredens logger */

   // logging to json with one argument
   50000	     25670 ns/op
   
   // logging to json without arguments
   50000	     42466 ns/op
```

## Running
```
make bench

// or simply
go test -bench .

```

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
