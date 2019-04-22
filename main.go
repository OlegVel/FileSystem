package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func directoryCheck(path string, depth int) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		logrus.WithError(err)
	}

	var b bytes.Buffer

	for i := 0; i < depth; i++ {
		b.WriteString("| ")
	}

	prefix := "├───"
	count := len(files) - 1

	for i, file := range files {
		if i == count {
			prefix = "└───"
		}
		if file.IsDir() {

			fmt.Println(b.String() + prefix + file.Name())
			directoryCheck(path+"/"+file.Name(), depth+1)
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

	directoryCheck(path, 0)
}
