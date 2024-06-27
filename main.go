package main

import (
	"fmt"

	"github.com/friendly-fhir/go-template/internal/greeting"
)

func main() {
	fmt.Println(greeting.New())
}
