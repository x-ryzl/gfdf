package main

import (
	"gfdf/common"
	"flag"
	"fmt"
)

var src = ""
var target = ""
var excludeDir = ""

func init()  {
	flag.StringVar(&src, "s", "", "源路径")
	flag.StringVar(&target, "t", "", "目标路径")
	flag.StringVar(&excludeDir, "e", "", "忽略文件夹")
	flag.Parse()
}
func main()  {
	if src == "" {
		fmt.Println("src dir required")
		panic("src dir required")
	}
	if target == "" {
		fmt.Println("target dir required")
		panic("target dir required")
	}
	filelist := common.FindDirFiles(src,src,excludeDir)
	filelist2 := common.FindDirFiles(target,target,excludeDir)
	difflist := common.AllDifference(filelist,filelist2)
	for _,f:= range difflist {
		fmt.Println("f:",f)
	}
	fmt.Println("len1:",len(filelist))
	fmt.Println("len2:",len(filelist2))

	fmt.Println("difflen:",len(difflist))

}
