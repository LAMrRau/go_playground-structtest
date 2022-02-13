# go_playground-structtest
Repo to test on structs, memory allocation and benchmarks.

Based on medium article: https://medium.com/better-programming/how-to-speed-up-your-struct-in-golang-76b846209587

Run the following command in the root directory of the repo to execute the benchmark
```go
go test -bench=. -benchtime=5s
```

The output should look similar to the following:

    [...]
    BenchmarkCreateGoodStructArray-4           69906             83616 ns/op
    BenchmarkCreateBadStructArray-4            53503            109083 ns/op
    BenchmarkTraverseGoodStruct-4            1757775              3413 ns/op
    BenchmarkTraverseBadStruct-4             1748144              3415 ns/op
    [...]

Run the following command in the root directory of the repo to execute the main function
```go
go run .
```

The output should look similar to the following:

    [...]
    good struct Type: main.GoodStruct Size:16
    bad struct Type: main.BadStruct Size:24

## Take-away message
* Go does not optimize the order of properties within a struct
* Go tries to put individual properties of a struct into 64 bit memory sections. It opens up a new 64 bit memory section if the next property does not fit the current on anymore.
* "Bad" ordering will result in more wasted memory
* "Bad" ordering will have impact on struct initialization
* "Bad" ordering has no impact on accessing properties afterwards
