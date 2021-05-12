package main

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// v1 - just embed a file into a string and print it
//
// NOTE: If you don't actually use the embed package, you need to import it
// with an underscore so that the compiler knows to look for the go:embed
// directive
var (
	//go:embed files/hello.gohtml
	hello string
)

func v1() {
	fmt.Println("V1")
	fmt.Print(hello)
}

// v2 - embed a directory and open a known file
var (
	//go:embed files/*
	embedFS embed.FS
)

func v2() error {
	fmt.Println("V2")
	f, err := embedFS.Open("files/hello.gohtml")
	if err != nil {
		return err
	}
	_, err = io.Copy(os.Stdout, f)
	return err
}

// v3 - iterate over files
func v3() error {
	fmt.Println("V3")
	entries, err := embedFS.ReadDir("files")
	if err != nil {
		return err
	}
	for _, entry := range entries {
		fmt.Println(entry.Name())
		f, err := embedFS.Open(filepath.Join("files", entry.Name()))
		if err != nil {
			return err
		}
		io.Copy(os.Stdout, f)
	}
	return nil
}

func main() {
	v1()
	err := v2()
	if err != nil {
		panic(err)
	}
	err = v3()
	if err != nil {
		panic(err)
	}
}
