package main

import (
	"flag"
	"fmt"
	"gfdf/comparison"
	"log"
)

var src = ""
var target = ""
var excludeDir = ""
var cmd = 0

const (
	dirdifffiles = iota
	samefile
)

func init() {
	flag.StringVar(&src, "s", "", "源路径")
	flag.StringVar(&target, "t", "", "目标路径")
	flag.StringVar(&excludeDir, "e", "", "忽略文件夹")
	flag.IntVar(&cmd, "c", 0, "命令")
	flag.Parse()
}
func main() {
	fmt.Println("cmd:", cmd)

	switch {
	case cmd == dirdifffiles:
		compareDirFiles()
	case cmd == samefile:
		findSameFile()

	}

}

func compareDirFiles() {
	if src == "" {
		fmt.Println("源文件必传")
		panic("源文件必传")
	}
	if target == "" {
		fmt.Println("目标文件必传")
		panic("目标文件必传")
	}

	var c util.Comparison
	c.Srcdir = src
	c.Targetdir = target
	c.Extdir = excludeDir

	difflist := c.DiffFolderFiles()
	for _, f := range difflist {
		fmt.Println("文件:", f)
	}

	fmt.Println("文件数量差:", len(difflist))
}

func findSameFile() {
	log.Println(src)
	log.Println(target)
	if src == "" {
		fmt.Println("源文件夹必传")
		panic("源文件夹必传")
	}
	if target == "" {
		fmt.Println("目标文件夹必传")
		panic("目标文件夹必传")
	}

	var c util.Comparison
	c.Srcfile = src
	c.Targetfile = target

	if c.SameFile() {
		fmt.Println(fmt.Sprintf("%s和%s为相同文件", src, target))
	} else {
		fmt.Println(fmt.Sprintf("%s和%s为不相同文件", src, target))
	}

}
