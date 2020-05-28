package main

import "gfdf/common"

type comparison struct {
	srcfile    string
	targetfile string
	srcdir     string
	targetdir  string
	extdir     string
}

func (c *comparison) SameFile() bool {
	md5v1, error := common.HashFileMd5(c.srcfile)
	md5v2, error2 := common.HashFileMd5(c.srcfile)

	if error != nil || error2 != nil {
		return false
	}

	if md5v1 == md5v2 {
		return true
	} else {
		return false
	}
}

func (c *comparison) DiffFolderFiles() []string {
	srcFolderFiles := common.FindDirFiles(c.srcdir, c.srcdir, c.extdir)
	targetFolderFiles := common.FindDirFiles(c.targetdir, c.targetdir, c.extdir)
	return common.AllDifference(srcFolderFiles, targetFolderFiles)
}
