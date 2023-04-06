package main

import (
	"archive/zip"
	"fmt"
	"path/filepath"
	"testing"
)

func Test(t *testing.T) {
	abs, err := filepath.Abs("../../fdsa")
	if err != nil {

	}
	fmt.Println(abs)
}

func Test2(t *testing.T) {
	join := filepath.Join("f/dsa", "ff")

	fmt.Println(join)

	reader, _ := zip.OpenReader("C:\\Users\\kks1234\\Desktop\\12.zip")

	for _, f := range reader.File {
		fmt.Println(f.Name)
	}

}

func Test3(t *testing.T) {
	ff()
	fmt.Println("hello")
}

func ff() {
	defer func() {
		fmt.Println("enter defer")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	nums := []int{1, 2, 3}
	var index = 4
	fmt.Println(nums[index])
	fmt.Println("done")

}
