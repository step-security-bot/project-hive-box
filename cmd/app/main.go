package main

import (
	"fmt"

	"github.com/Jorik-VanLooy/project-hive-box/handlers"
)

func main() {
	str := handlers.PrintHello()
	fmt.Println(str)
}
