package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var separator = string(os.PathSeparator)

//遍历获取获取文件夹下的所有文件
func FindDirFiles(src, base, exclude string) (filelist []string) {
	files, _ := ioutil.ReadDir(src)
	for _, onefile := range files {
		if onefile.IsDir() {
			if strings.Contains(exclude, onefile.Name()) {
				continue
			}
			filelist = append(filelist, FindDirFiles(src+separator+onefile.Name(), base, exclude)...)
		} else {
			filename := strings.Replace(fmt.Sprintf("%s%s%s", src, separator, onefile.Name()), base, "", -1)
			filelist = append(filelist, filename)
			//fmt.Println("src: ",fmt.Sprintf("%s/%s", src, onefile.Name()))
		}
	}
	return
}

//求并集
func Union(slice1, slice2 []string) []string {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

//求交集
func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

//求差集 slice1-并集
func Difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

//求全差集 slice1+slice2-并集
func AllDifference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	allslice := append(slice1, slice2...)

	for _, value := range allslice {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

// 获取文件的md5值
func HashFileMd5(filePath string) (md5Str string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return
	}
	hashInBytes := hash.Sum(nil)[:16]
	md5Str = hex.EncodeToString(hashInBytes)
	return
}
