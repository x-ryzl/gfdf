package main

import (
	"flag"
	"fmt"
)

var src = ""
var target = ""
var excludeDir = ""

func init() {
	flag.StringVar(&src, "s", "", "源路径")
	flag.StringVar(&target, "t", "", "目标路径")
	flag.StringVar(&excludeDir, "e", "", "忽略文件夹")
	flag.Parse()
}
func main() {
	if src == "" {
		fmt.Println("src dir required")
		panic("src dir required")
	}
	if target == "" {
		fmt.Println("target dir required")
		panic("target dir required")
	}

	var c comparison
	c.srcdir = src
	c.targetdir = target
	c.extdir = excludeDir

	difflist := c.DiffFolderFiles()
	for _, f := range difflist {
		fmt.Println("f:", f)
	}

	fmt.Println("difflen:", len(difflist))

}
