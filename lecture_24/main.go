package main

import (
	"fmt"
	"os"
)

func main() {

	// dir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(dir)

	// isErr := createFile("master_academy2.txt", "hello bangladesh")
	// fmt.Println(isErr)

	fi, err := os.Stat("master_academy2.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(fi.IsDir())
	fmt.Println(fi.ModTime().Date())
	fmt.Println(fi.Name())
	fmt.Println(fi.Size())

}

func createFile(fileName, content string) bool {

	posf, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	defer posf.Close()
	_, err = posf.Write([]byte(content))
	//fmt.Println(n, err)
	if err != nil {
		return false
	}
	return true
}
