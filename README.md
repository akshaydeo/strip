## Strip

A utility to comment/uncomment the un-necessary function calls during the builds.

> This is in development

> From the results, it does not seem like this lib could become useful :-/

> Just keeping it as is for reference 

## Example

- During the development, we keep on adding a lot of logging statements, some of which involve allocations/reflection/resource intensive function calls. 
- During the live load, we do not need them to execute at all, instead of just skipping console logging.

```go
logger.D(reflect.TypeOf(i).String())
``` 

You can use Strip as a part of the build system, to skip the calls like `logger.D` during the final builds.


## Usage

### Command line

```bash
strip -pkg log -call log.Println -path demo/ -v -r
``` 

| Param | Meaning| 
|---|---|
|-pkg|Package name to consider|
|-call|Calls to process|
|-r|Recursive|
|-u | Revert the changes done by strip call|
|-v|Verbose logging| 



## Tests with regular logging

```go
// Function to log the interface passed to it
// I am using logrus as an example, with different levels
func logIt(i interface{}) {
	t, _ := json.Marshal(i)
	logrus.Debug(string(t))
}
```
> i is a struct with sufficient amount of complexity for the context


### With logrus
- DebugLevel: 1940 ms
- ErrorLevel: 1657 ms
- Stripped: 1621 ms

36 ms saved against 10k operations :-|


#### Details

- Debug
```bash
goos: darwin
goarch: amd64
pkg: github.com/akshaydeo/strip
BenchmarkStripping-12    	   10000	    186498 ns/op	   38564 B/op	     494 allocs/op
```
- Error
```bash
goos: darwin
goarch: amd64
pkg: github.com/akshaydeo/strip
BenchmarkStripping-12    	   10000	    160190 ns/op	   32521 B/op	     418 allocs/op
PASS

```
- Stripped
```bash
goos: darwin
goarch: amd64
pkg: github.com/akshaydeo/strip
BenchmarkStripping-12    	   10000	    160179 ns/op	   32409 B/op	     417 allocs/op
PASS
```

### Tests with JSON logging

```go
// Function to log the interface passed to it
// I am using logrus as an example, with different levels
func logIt(i interface{}) {
	t, _ := json.Marshal(i)
	logrus.Debug(string(t))
}
```

- Debug
```bash
goos: darwin
goarch: amd64
pkg: github.com/akshaydeo/strip
BenchmarkStripping-12    	   10000	    197467 ns/op	   42860 B/op	     458 allocs/op
PASS
```

- Error
```bash
goos: darwin
goarch: amd64
pkg: github.com/akshaydeo/strip
BenchmarkStripping-12    	   10000	    172496 ns/op	   36165 B/op	     441 allocs/op
PASS
```

- Stripped
```bash
goos: darwin
goarch: amd64
pkg: github.com/akshaydeo/strip
BenchmarkStripping-12    	   10000	    165468 ns/op	   32408 B/op	     417 allocs/op
PASS

```