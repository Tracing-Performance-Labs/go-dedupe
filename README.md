# go-dedupe

This is a simple library for deduplicating strings in Go.

## Installation
```bash
go get github.com/Tracing-Performance-Labs/go-dedupe
```

## Sample usage

```go
package main

import (
    "fmt"
    "github.com/Tracing-Performance-Labs/go-dedupe"
)

func main() {
    codec := dedupe.NewCodec()
    compacted := codec.Encode("hello")
    fmt.Println("Compacted:", compacted)
}
```
    
You can also run the bundled executables to test the library. 

## License
MIT
