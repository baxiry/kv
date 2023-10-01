# key value in memory store

### safe concurrent key value in memory store

## install:
```
go get github.com/baxiry/kv 
```

## usage:
```go
package main

import (
      "github.com/baxiry/kv"
)

func main() {
      imap := new(kv.Map[int, int])

      imap.Set(1, 111) // insert
      imap.Set(1, 123) // update

      imap.Set(2, 222)

      val, err := imap.Get(1) //  123
      val, err = imap.Get(10) //  not found error
 
      ok := imap.HasKey(2) //  true
      ok = imap.HasKey(20) //  false

      imap.Delete(2)

      ok = imap.HasKey(2) // false

      // new stringer Map 
      strMap := new(kv.Map[string, string])

      strMap.Set("hi", "hello")
}
```

### TODO:

- [x] Set
- [x] Get
- [x] HasKey
- [x] Delete
- [x] full testing.
### not important
- [ ] auto delete by timeout.
- [ ] benchmark.
- [ ] avoid pointers as mach as posible for more gc effecion.

### license:

``` 
use it with any license you prefer
```
