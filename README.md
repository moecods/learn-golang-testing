# Tutorial golang testing

When i work my own golang projects, i undrestand i cant learn golang testing if i dont focus on it
so i create this repo for learn about testing in golang

Im very happy to start learning golang but if i cant write test for my projects after a long time i cant
improve my golang projects and my skills

the `testing` package in Go provides the tools necessary to write unit tests.

***
for first try first i create math file and Add func
Add func get two int value and return sum and return sum of them
i created a file test with name math_test (math + _test)
i write test for Add func after that for running my test i run statement go test**
***
you can run test by run below statement:
```
go test
```
***

The `*testing.T` type provides methods like `Error`, `Errorf`, `Fatal`, `Fatalf` Witch are used to report failture during a test.

- Error: Marks the test as failed but continues execution.
- Fatal: Marks the test as failed and stops execution immediately.
***
I write Table Driven Test for Add function in meth_test
table-driven test are a pattern used in Go to test multiple scenarios with diffrent inputs and expected outputs.
***
check coverage: go test -cover
run the benchmark: go test -bench=.