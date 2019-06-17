package main

import (
	"fmt"
	"math"

	"github.com/envacance/go_crash_course/03_packages/strutil" // import
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))

	fmt.Println(strutil.Reverse("olleh"))

}
