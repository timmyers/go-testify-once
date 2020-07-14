```
❯ go test -test.v ./...
?       github.com/timmyers/go-testify-once     [no test files]
=== RUN   TestSuiteA
=== RUN   TestSuiteA/TestDoA
Only ONCE
Do A
--- PASS: TestSuiteA (0.00s)
    --- PASS: TestSuiteA/TestDoA (0.00s)
PASS
ok      github.com/timmyers/go-testify-once/pkg/a       (cached)
=== RUN   TestSuiteB
=== RUN   TestSuiteB/TestDoB
Only ONCE
Do B
--- PASS: TestSuiteB (0.00s)
    --- PASS: TestSuiteB/TestDoB (0.00s)
PASS
ok      github.com/timmyers/go-testify-once/pkg/b       (cached)
?       github.com/timmyers/go-testify-once/pkg/test    [no test files]
```
