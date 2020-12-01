# int2bin
import 	"github.com/andvary/int2bin"

int2bin produces binary representations of integers (int, int8/16/32/64, uint, uint8/16/32/64), which may be useful, if 
you deal with binary arithmetics a lot. 

Example:
```
package main

import (
    "fmt"
    "log"

    "github.com/andvary/int2bin"
)

func main() {
    binString, err := int2bin.Bin(-1000)
    if err != nil {
	    log.Fatal(err)
    }

    fmt.Println(binString)
}
```

Produces:

```1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1100 0001 1000```


<b>Why not fmt.Printf("%b", n)?
1. Go stores negative numbers as 2's complements, whereas %b produces the regular representation with negative sign.
2. %b displays strings of bits without spaces making long sequences somewhat hard to read.

