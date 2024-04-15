Benchmarks of various methods to add css classes to a pre-existing html string,
for a philosophical reason.

```sh
❯ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/tts2k/resep-template-bench
cpu: AMD Ryzen 7 6800H with Radeon Graphics         
BenchmarkPango2-16              1000000000          0.003205 ns/op
BenchmarkGoquery-16             1000000000          0.001620 ns/op
BenchmarkGoTemplate-16          1000000000          0.0008292 ns/op
BenchmarkMustache-16            1000000000          0.0004077 ns/op
BenchmarkFastTemplate-16        1000000000          0.0001359 ns/op
PASS
ok      github.com/tts2k/resep-template-bench   0.065s

```


