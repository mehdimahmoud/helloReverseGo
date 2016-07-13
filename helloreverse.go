/*
 * Example of how to use our own built Library
 * Prints the reverse string
 */
package main

import (
	"fmt"

	"github.com/mehdimahmoud/stringutilGo"
)

func main() {
	fmt.Println(stringutil.Reverse("! oG, olleH"))
	fmt.Println(stringutil.Reverse("Hello, Go!"))
}
