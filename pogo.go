package main

import (
	"crypto/md5"
	"hash/crc32"
	"encoding/hex"
	"strconv"
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

		c := getConfig()

		var destSum, sourceSum string

		if c.hashType == "md5" {
			destSum = checkMd5(sf)
			sourceSum = checkMd5(df)
		}

		if c.hashType == "crc32" {
			destSum = checkCrc32(source)
			sourceSum = checkCrc32(dest)
		}

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


func checkMd5(file io.Reader) (sum string) {
	h := md5.New()
	io.Copy(h, file)
	sum = hex.EncodeToString(h.Sum(nil))
	return sum
}

func checkCrc32(fileName string) (sum string) {
	h := crc32.NewIEEE()
	b, err := ioutil.ReadFile(fileName)

	if err != nil {
		printErrAndExit("(crc32) Error reading: " + fileName)
	} else {
		h.Write(b)
		sum = strconv.FormatUint(uint64(h.Sum32()), 10)
		return sum
	}

	return
}
