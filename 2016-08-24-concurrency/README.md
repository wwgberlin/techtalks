# WWG Berlin 2016-08-24 @ [Talon.One] - Concurrency

This is a bit of code extracted from our production server tests, used to
demonstrate using Go routines for concurrent tasks. In [main.go][] you will find
a dummy HTTP service, and in [main_test.go][] a test that performs a series of
requests against the server. The commit history of the repo shows taking this
test runner from serial to concurrently running tests.

[Talon.One]: http://talon.one
[main.go]: main.go
[main_test.go]: main_test.go

## The "demo"

1. Get the code: `go get github.com/wwgberlin/techtalks/2016-08-24-concurrency`.
2. Change to the package dir: `cd $GOPATH/src/github.com/wwgberlin/techtalks/2016-08-24-concurrency`.
3. Fetch git tags for easier navigation: `git fetch`.
4. Check out the "before" version `git checkout step-0`.
5. Run the tests: `go test`. Observe the time taken.
6. Check out the "after" version `git checkout master`.
7. Run the tests again: `go test`. Observe the speedup.
8. Inspect what changed with `git show step-1`, `git show step-2` etc. 

## License

```
Copyright (c) 2016 Talon.One GmbH

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```
