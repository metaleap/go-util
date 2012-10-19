package main

import (
	"fmt"
	"os"
	"path/filepath"

	ioutil "github.com/go-ngine/go-util/io"
)

func main() {
	fmt.Println(os.Getenv("GOPATH"))
	return
	ioutil.WalkDirectory(filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "go-ngine"), ".go")
}
