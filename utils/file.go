package utils

import (
	"io/ioutil"
	"log"
	"strings"
)

func ReplaceData(path string, originalString string, newString string) (err error) {
	if fileData, err := ioutil.ReadFile(path); err != nil {
		log.Fatal(err)
	} else {
		newContent := []byte(strings.Replace(string(fileData), originalString, newString, -1))
		err = ioutil.WriteFile(path, newContent, 0)
	}
	return err
}
