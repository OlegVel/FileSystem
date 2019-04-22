package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func directoryCheck(path, prefix string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		logrus.WithError(err)
	}

	count := len(files) - 1

	for i, file := range files {
		if i == count {
			//prefix = " └" + prefix + "─"
		}
		if file.IsDir() {
			fmt.Println(prefix + file.Name())
			directoryCheck(path+"/"+file.Name(), "|  "+prefix+"─")
		} else {
			fmt.Println(prefix + file.Name() + " (" + fmt.Sprint(file.Size()) + "b)")
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

	directoryCheck(path, "├──")
}
