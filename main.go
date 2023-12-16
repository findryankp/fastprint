package main

import (
	"fmt"

	"github.com/Findryankp/mvcGo"
)

func main() {

	if err := mvcGo.Init(); err != nil {
		fmt.Println(err)
		return
	}

	//run on port 8080
	LoadConfig(":8080")

}
