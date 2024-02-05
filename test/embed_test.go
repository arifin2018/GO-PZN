package test

import (
	"embed"
	"fmt"
	_ "image/png"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T)  {
	fmt.Println(version)
}

//go:generate ../images/nikel.png ./local-asset-dir
//go:embed local-asset-dir
var nikel []byte
func TestByte(t *testing.T)  {
	err := ioutil.WriteFile("logonew.png",nikel, fs.ModePerm)
	if err != nil{
		panic(err)
	}
}

func TestWriteFile(t *testing.T)  {
	alias := "./images/nikel.png"
	idx := strings.Index(alias, filepath.Ext(alias)); 
	filename := fmt.Sprintf("%s%s", alias[:idx],filepath.Ext(alias))
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fileLocation := filepath.Join(dir, "tes", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer targetFile.Close()
	file, err := os.Open(alias)
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(targetFile,file); err != nil {
		panic(err)
	}
	fmt.Println(fileLocation)
}

//go:generate ../files/a.txt ./local-asset-dir
//go:generate ../files/b.txt ./local-asset-dir
//go:generate ../files/c.txt ./local-asset-dir
//go:embed local-asset-dir
var files embed.FS
func TestMultipleFIles(t *testing.T)  {
	file, err := files.ReadFile("files/a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))

	fileB, err := files.ReadFile("files/b.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fileB))

	fileC, err := files.ReadFile("files/c.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fileC))
}

//go:generate ../files ./local-asset-dir
//go:embed local-asset-dir
var path embed.FS
func TestPathMatcher(t *testing.T){
	dir, _ :=path.ReadDir("files")
	for k, v := range dir {
		if !v.IsDir() {
			fmt.Println(k,v.Name())
			content, _ := path.ReadFile("files/"+v.Name())
			fmt.Println("content => ", string(content))
		}
	}
}