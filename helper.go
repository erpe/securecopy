package main

import (
	"fmt"
	"os"
)

func checkSrc(dir string) {
	var ret, _ = exists(dir)
	if ret == false {
		message := "No such sourcedir: " + dir
		printErrAndExit(message)
	}
}

func checkDst(dir string) {
	var ret, _ = exists(dir)
	if ret != false {
		message := "Destination directory already exists: " + dir
		printErrAndExit(message)
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return false, err
}

func printErrAndExit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
