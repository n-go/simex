// +build ignore

package main

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	generate("./")
}

func generate(dir string) {
	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, f := range fs {
		if f.IsDir() {
			generate(filepath.Join(dir, f.Name()))
			continue
		}
		if !strings.HasSuffix(f.Name(), ".json") {
			continue
		}
		out, err := os.Create(path.Join(dir, f.Name()+"_test.go"))
		if err != nil {
			panic(err)
		}
		out.WriteString("package simex_test\n\nconst ")
		out.WriteString(strings.TrimSuffix(f.Name(), ".json"))
		out.WriteString("JSON = `\n")
		file, _ := os.Open(path.Join(dir, f.Name()))
		io.Copy(out, file)
		out.WriteString("`\n")
	}
}
