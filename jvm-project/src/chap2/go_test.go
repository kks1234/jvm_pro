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
