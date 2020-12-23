package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	rn := os.Getenv("RONAN_CATALOG_USER")

	os.Setenv("FOO", "1")
	fmt.Println("Ronan:", rn)
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
