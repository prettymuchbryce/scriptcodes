#### Bitcoin Script Opcodes

All Bitcoin Script opcodes. For use in your bitcoin-related go projects.

Taken from the [Bitcoin Wiki](https://en.bitcoin.it/wiki/Script)

Install:
```
go get github.com/prettymuchbryce/scriptcodes
```

Usage:
```
    package main

    import (
        "fmt"
        "scriptcodes"
    )

    func main() {
        fmt.Println(scriptcodes.OP_VERIFY)
    }
```

Enjoy!
