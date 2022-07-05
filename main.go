package main

import (
	"fmt"

	"github.com/lans89/ConverterMeasure/distances"
)

func main() {
	fmt.Println("Test Module")
	fmt.Println(distances.ConvertDistance("m", "Km", 3456))
}
