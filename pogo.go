package main

import (
	"fmt"
	"io/ioutil"
)

type SourceFiles struct {
	path
}

func (source *SourceFiles) RecursiveCopy(topath string) {
	os.MkdirAll(topath, 0777)
	var list, err = ioutil.ReadDir(source.path)
	if err == nil {
		for _, val := range list {
			fmt.Println("item: ", val.Name())
		}
	} else {
		fmt.Println("Error: ", err)
	}
}


func CopyDir(from string, to string) (err error) {
}

func CopyFile(from string, to string) (err error) {
	sf, err := os.Open(from)
	if err != nil {
		return err
	}
	defer sf.Close()
	df, err = os.Create(to)
}
