package main

import (
	"flag"
	c "./MakeDirsAndFiles"
	"fmt"
	"log"
	fp "path/filepath"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "f", "", "file full path")
	flag.Parse()
	
	fmt.Println(filePath)
		

	dir, _ := fp.Split(filePath)
	c.CreateDirIfNotExist(dir)
	err := c.CreateFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
}
