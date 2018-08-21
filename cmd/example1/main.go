package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("example1 env FOO=%s\n", os.Getenv("FOO"))
}
