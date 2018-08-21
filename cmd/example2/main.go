package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("example2 env FOO=%s\n", os.Getenv("FOO"))
}
