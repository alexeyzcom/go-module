package main

import (
	"fmt"
	"github.com/alexeyzcom/go-module/internal/app"
	"github.com/alexeyzcom/mathmod"
)

func main() {
	foo := mathmod.SumWithLog(15, 20)

	fmt.Println(foo)
	app.Start()
}
