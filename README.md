GoSort - Benchmarking Go's default sort
=======================================

To get the code and run the benchmarks sorting 1000 numbers:

    $ git clone https://github.com/absoludity/gosort.git && cd gosort && go test -bench=".*1000"


Example results (on cheap netbook hardware)
-------------------------------------------
    $ git clone https://github.com/absoludity/gosort.git && cd gosort && go test -bench=".*1000"
    PASS
    BenchmarkDefaultSort1000            1000           1515159 ns/op
    BenchmarkSort1000            100          13021620 ns/op
    BenchmarkSort1000WithInterface        20          82104350 ns/op
