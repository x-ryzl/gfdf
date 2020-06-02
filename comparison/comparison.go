package util

import (
	"fmt"
	"gfdf/common"
)

type Comparison struct {
	Srcfile    string
	Targetfile string
	Srcdir     string
	Targetdir  string
	Extdir     string
}

//根据文件的md5值判读是否为同一文件
func (c *Comparison) SameFile() bool {
	fmt.Println(c.Srcfile)
	md5v1, error := common.HashFileMd5(c.Srcfile)
	md5v2, error2 := common.HashFileMd5(c.Targetfile)
	fmt.Println(md5v1)
	fmt.Println(md5v2)
	if error != nil || error2 != nil {
		return false
	}

	if md5v1 == md5v2 {
		return true
	} else {
		return false
	}
}

//获取两个文件路径下的文件不一样的文件列表(根据名称来判断)
//extdir: 忽略的文件夹
func (c *Comparison) DiffFolderFiles() []string {
	srcFolderFiles := common.FindDirFiles(c.Srcdir, c.Srcdir, c.Extdir)
	targetFolderFiles := common.FindDirFiles(c.Targetdir, c.Targetdir, c.Extdir)
	return common.AllDifference(srcFolderFiles, targetFolderFiles)
}
