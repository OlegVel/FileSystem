package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func directoryCheck(path string, depth []bool) {
	var b bytes.Buffer
	var prefix string
	var depthIn []bool

	files, err := ioutil.ReadDir(path)
	if err != nil {
		logrus.WithError(err)
	}

	for _, depthFlag := range depth {
		if depthFlag {
			b.WriteString(" \t")
		} else {
			b.WriteString("|\t")
		}

	}

	count := len(files) - 1

	for i, file := range files {
		if i == count {
			prefix = "└────"
			depthIn = append(depth, true)
		} else {
			prefix = "├────"
			depthIn = append(depth, false)
		}

		if file.IsDir() {
			fmt.Println(b.String() + prefix + file.Name())
			directoryCheck(path+"/"+file.Name(), depthIn)
		} else {
			fmt.Println(b.String() + prefix + file.Name() + " (" + fmt.Sprint(file.Size()) + "b)")
		}
	}
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		logrus.WithError(err)
		return
	}
	fmt.Println(path)

	directoryCheck(path, []bool{})
}
