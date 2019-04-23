package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
)

func directoryCheck(writer io.Writer, path string, depth []bool) {
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
			fmt.Fprintln(writer, b.String()+prefix+file.Name())
			directoryCheck(writer, path+"/"+file.Name(), depthIn)
		} else {
			fmt.Fprintln(writer, b.String()+prefix+file.Name()+" ("+fmt.Sprint(file.Size())+"b)")
		}
	}
}

func main() {
	writer, _ := os.Create("./path.txt")
	//path, err := os.Getwd()
	//if err != nil {
	//	logrus.WithError(err)
	//	return
	//}

	path := "/home/"
	fmt.Println(path)

	directoryCheck(writer, path, []bool{})
}
