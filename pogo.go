package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"securecopy/protocoll"
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

	cfg := getConfig()

	if (dest != cfg.destinationDir) && (!os.IsNotExist(err)) {
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

	defer sf.Close()

	df, err := os.Create(dest)

	if err != nil {
	  return err
	}

	defer df.Close()

	_, err = io.Copy(df, sf)
	if err == nil {
	  si, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, si.Mode())
		}
		checkHash(sf)
		destSum := CheckMd5(df)
		sourceSum := CheckMd5(sf)
		if destSum == sourceSum {
			go protocoll.Success(dest + " : " + sourceSum + " : GOOD" )
			go fmt.Print("+")
		} else {
			go protocoll.Failure(source + " : " + sourceSum + " : MISMATCH")
			go fmt.Print("E")
		}
	}
	return
}

func checkHash(file io.Reader) (sum string) {
	c := getConfig()
	fmt.Println("config: " + c.hashType)
	return "foobar"
}

func CheckMd5(file io.Reader) (sum string) {
	md5 := md5.New()
	io.Copy(md5, file)
	sum = hex.EncodeToString(md5.Sum(nil))
	return sum
}

