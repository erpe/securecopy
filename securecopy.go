package main

import (
	"flag"
	"fmt"
	"github.com/erpe/securecopy/protocoll"
	"io/ioutil"
	"os"
)

var sDir, sourceDir string
var fileMap map[string]string
var cfg Config

type Config struct {
	destinationDir string
}

func main() {
	srcDir := flag.String("source", "", "the source directory")
	destDir := flag.String("destination", "", "the to be created destination directory")

	flag.Parse()
	if flag.NFlag() < 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	makeConfig(*destDir)
	fileMap = make(map[string]string)

	checkSrc(*srcDir)
	checkDst(*destDir)
	protocoll.Initialize(*destDir)
	err := CopyDir(*srcDir, *destDir)
	if err != nil {
		fmt.Println("\nError while copying: ", err)
	} else {
		fmt.Println("\nfinished copying: ", *srcDir)
		fmt.Println("see : ", *destDir+"/protocol.txt for details...")
	}
}

func makeConfig(str string) {
	cfg = Config{str}
}

func getConfig() (ret Config) {
	ret = cfg
	return ret
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
