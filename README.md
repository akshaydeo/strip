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


## Usage

### Command line

```bash
strip --pkg pname --call pname.D --call pname.E -r --vendor  
``` 

| Param | Meaning| 
|---|---|
|--pkg|Package name to consider|
|--call|Calls to process|
|-r|Recursive|
|--unstrip | Revert | 

