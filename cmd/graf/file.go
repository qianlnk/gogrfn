package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	graf "github.com/qianlnk/graf"
)

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func save(data []byte, name string) error {
	seoName := graf.SEOURL(name)
	fnames := strings.Split(seoName, ".")
	var filename string
	if fnames[len(fnames)-1] != "json" {
		filename = fmt.Sprintf("%s.json", seoName)
	} else {
		filename = seoName
	}

	if checkFileIsExist(filename) {
		return errors.New(fmt.Sprintf("name: %s exist already", name))
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	return err
}
