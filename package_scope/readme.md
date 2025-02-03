### Package Scope

Package scope mainly works with packages from other folders. Like we are creating multiple folder in lib folder. now to access from other folders we need to use Package scope. 

⇒ best practice for declaring package name is .. naming it with folder name

⇒ to create go.mod we should call 'go mod init example.com'

```go
package main

// importing package from another folder
// to create go.mod we should call 'go mod init example.com'
import (
	"example.com/mathlib"
)

func main() {
	mathlib.Add(4, 5)
}

```

```go
package mathlib

import "fmt"

func Add(n1 int, n2 int) {
	res := n1 + n2
	fmt.Println(res)
}

```

here we have math.go in mathlib folder. now to access math.go from main.go we need to create go.mod

```go
go mod init example.com
```

we need to write this to create go.mod. we can give any name instead of exapmle.com