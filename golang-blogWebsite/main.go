package main

import (
	"blog/service"
	"fmt"
)

func main() {
	sanjeev, err := service.NewUser(1, "Sanjeev", "sanjeevyadav368@gmail.com", "123455")
	if err != nil {
		fmt.Println("There is some error", err.Error())
	}
	fmt.Println("the user created is :", sanjeev)
}
