## Strip

A utility to comment/uncomment the un-necessary function calls during the builds.

> Yet to complete :-)

## Example

- During the development, we keep on adding a lot of logging statements, some of which involve allocations/reflection/resource intensive function calls. 
- During the live load, we do not need them to execute at all, instead of just skipping console logging.

```go
logger.D(reflect.TypeOf(i).String())
``` 

You can use Strip as a part of the build system, to skip the calls like `logger.D` during the final builds.


## Configuration
`Stripe` searches for `config.strip`, which is a yaml file, and performs the operations accordingly. If the file is not present, it will panic and will be a NOP.


### Sample config file

```yaml
stripe:
  version: 1
  stripe:
    call:
      - logger.Print
      - log.Println
  build:
    - go build main.go
```
- This flow will, comment out all the function call references of `logger.Print` and `log.Println`
- Then it will execute, commands listed under build
- And at the end will revert the code back, i.e. uncomment the commented function calls.