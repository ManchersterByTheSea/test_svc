package main

import (
	"fmt"
	"github.com/ManchersterByTheSea/test_common/register"
	_ "github.com/ManchersterByTheSea/test_drive"
)

func main() {
	driver := register.Drivers["test"]
	driver.Create()
	err := driver.Update("hhh", 1)
	fmt.Println(err)
}
