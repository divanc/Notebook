package main

import (
	"fmt"
	"math"

	"github.com/Betra/go_crash/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(16))
	fmt.Println(strutil.Reverse("Hello"))
}
