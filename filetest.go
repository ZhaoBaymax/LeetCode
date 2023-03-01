package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/null")
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("close error: ", err)
		} else {
			fmt.Println("close no error")
		}
	}()
	if err != nil {
		fmt.Println("open error! ", err)
		return
	}

}
