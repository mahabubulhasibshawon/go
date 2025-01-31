## What is scope

scope declares the availability of variables like can we access any specific func is declared by “scope”

```go
// code
package main
import "fmt"

var a = 20
var b = 30

func add(x int, y int) {
	z := x + a
	fmt.Println(z)
}

func main () {
	p := 30
	q := 40
	
	add(p,q)
	
	add(a,b)
	
	// wrong
	add(a,z)
}
```

Here at the example we can see we can access global variables (a,b) from any function but we can not access local variable of add function from main function as z is not declared in main. here we don’t have any scope of z in main 

### Local Scope

variables declared in main func is local scope

we can create local bloc inside a local bloc, 

- checking criteria for local scope :
    - first check the block you’re in
    - if not there then if will check main.go
    - if it is not there then it will check global scope

```go
// code
package main
import "fmt"

var a = 20
var b = 30

func main () {
	p := 30
	q := 40
	if p > 25 { 
		x := "Should get married"
		fmt.Println("As age of Rahim(",p,"), he",x, "otherwise he may to get a wife by ",q)
		fmt.Println(z) // error
	}	
}

```

at the example we can see we have local scope p,q and global scope a, b

we also have local scope inside if () {} block as x.