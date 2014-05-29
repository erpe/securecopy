package main

import (
  "fmt"
	"flag"
	"io/ioutil"
	"os"
  )

var sDir, sourceDir, destDir string
func main() {
	srcDir := flag.String("source", "" , "the source directory")
	destDir := flag.String("destination", "", "the to be created destination directory")
	flag.Parse()
	if flag.NFlag() < 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	checkSrc(*srcDir)
	checkDst(*destDir)
	os.MkdirAll(*destDir, 0777)
	sourcefiles :=  SourceFiles{path: *srcDir}
	sourcefiles.Copy(*destDir)
	fmt.Println("Created new Directory: " + *destDir)
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
