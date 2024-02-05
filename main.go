package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed logonew.png
var logo []byte

//go:embed files/*.txt
var path embed.FS
func main()  {
	fmt.Println(version)

	err := ioutil.WriteFile("logonew.png",logo, fs.ModePerm)
	if err != nil{
		panic(err)
	}

	dir, _ :=path.ReadDir("files")
	for k, v := range dir {
		if !v.IsDir() {
			fmt.Println(k,v.Name())
			content, _ := path.ReadFile("files/"+v.Name())
			fmt.Println("content => ", string(content))
		}
	}
}