package main

// importing package from another folder
// to create go.mod we should call 'go mod init example.com'
import (
	"example.com/mathlib"
)

func main() {
	mathlib.Add(4, 5)
}
