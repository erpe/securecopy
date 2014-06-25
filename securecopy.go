package main

import (
  "fmt"
	"flag"
	"io/ioutil"
	"os"
	"securecopy/protocoll"
  )

var sDir, sourceDir, destDir string
var fileMap map[string]string

func main() {
	srcDir := flag.String("source", "" , "the source directory")
	destDir := flag.String("destination", "", "the to be created destination directory")
	//fileMap := make(map[string]string)
	flag.Parse()
	if flag.NFlag() < 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	fileMap = make(map[string]string)
	checkSrc(*srcDir)
	checkDst(*destDir)
	protocoll.Initialize(*destDir)
	err := CopyDir(*srcDir, *destDir)
	if err != nil {
		fmt.Println("Error while copying: ", err)
	} else {
		fmt.Println("copied: ", *srcDir)
	}
	fmt.Println(fileMap)
}


func listDirectory(dir string) {
	var list, err = ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		for _, val := range list {
			fmt.Println("item: ", val.Name())
		}
	}
}
