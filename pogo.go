package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type CopyError struct {
	What string
}

func (e *CopyError) Error() string {
  return e.What
}


func CopyDir(source string, dest string) (err error) {
	fi, err := os.Stat(source)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		msg := "Source is not a directory: " + source
		return &CopyError{ msg }
	}

	_, err = os.Open(dest)

	if !os.IsNotExist(err) {
		msg := "Destination already exists: " + dest
		return &CopyError{ msg }
	}

	err = os.MkdirAll(dest, fi.Mode())
	if err != nil {
		return err
	}

	entries, err := ioutil.ReadDir(source)

	for _, entry := range entries {
		sfp := source + "/" + entry.Name()
		dfp := dest + "/" + entry.Name()
		if entry.IsDir() {
			err = CopyDir(sfp, dfp)
			if err != nil {
				log.Println(err)
			}
		} else {
			err = CopyFile(sfp, dfp)
			if err != nil {
				log.Println(err)
			}
		}
	}
	return
}


func CopyFile(source string, dest string) (err error) {
  sf, err := os.Open(source)
	if err != nil {
	  return err
	}

	CheckMd5(sf)
	defer sf.Close()
	df, err := os.Create(dest)
	if err != nil {
	  return err
	}

	CheckMd5(df)
	defer df.Close()
	_, err = io.Copy(df, sf)
	if err == nil {
	  si, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, si.Mode())
		}
	}
	return
}

func CheckMd5(file io.Reader) {
	md5 := md5.New()
	io.Copy(md5, file)
	fmt.Println("md5: ")
	fmt.Printf("%x\t%s\n", md5.Sum(nil), file )
}

