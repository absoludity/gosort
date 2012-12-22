GoSort - Benchmarking Go's default sort
=======================================

To get the code and run the benchmarks sorting 1000 numbers:

$ git clone https://github.com/absoludity/gosort.git && cd gosort && go test -bench=".*1000"

Strangely, the ouput seems to show that a very simple bubble-sort is over
100 times faster than Go's [built-in sort][1]. I was unsure if this was a
result of the sort.Interface implementation (ie. extra function calls rather
than simple array indexing), so I also added a benchmark of the simple
bubblesort using sort.Interface, which is still over 15 times faster than Go's
built in sort.

Example results (on cheap netbook hardware)
-------------------------------------------
$ git clone https://github.com/absoludity/gosort.git && cd gosort && go test
-bench=".*1000"
PASS
BenchmarkDefaultSort1000            2000           1159938 ns/op
BenchmarkSort1000         200000              8917 ns/op
BenchmarkSort1000WithInterface     50000             69624 ns/op

[1] http://golang.org/pkg/sort/ "Go's sort package"


