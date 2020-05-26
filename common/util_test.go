package common

import (
	"fmt"
	"testing"
)

func TestIntersect(t *testing.T)  {
	var l1 []string
	var l2 []string

	l1 = append(l1,"2")
	l1 = append(l1,"1")
	l1 = append(l1,"3")

	l2 = append(l2,"2")
	l2 = append(l2,"3")
	l2 = append(l2,"4")

	if len(Intersect(l1,l2)) != 2 {
		t.Error("common.Intersect() error")
	}
}

func TestUnion(t *testing.T)  {
	var l1 []string
	var l2 []string

	l1 = append(l1,"2")
	l1 = append(l1,"1")
	l1 = append(l1,"3")

	l2 = append(l2,"2")
	l2 = append(l2,"3")
	l2 = append(l2,"4")

	if len(Union(l1,l2)) != 4 {
		t.Error("common.Intersect() error")
	}
}



func TestDiffence(t *testing.T)  {
	var l1  []string
	var l2  []string

	l1 = append(l1,"1")
	l1 = append(l1,"5")
	l1 = append(l1,"3")

	l2 = append(l2,"2")
	l2 = append(l2,"3")
	l2 = append(l2,"4")

	l3 := Difference(l1,l2)
	fmt.Println("l3:",l3)
	fmt.Println("l1:",l1)
	fmt.Println("l2:",l2)
	if len(l3) != 2 {
		t.Error("common.Difference() error,len: ",len(l3),"")
	}
}


func TestAllDiffence(t *testing.T)  {
	var l1  []string
	var l2  []string

	l1 = append(l1,"1")
	l1 = append(l1,"5")
	l1 = append(l1,"3")

	l2 = append(l2,"2")
	l2 = append(l2,"3")
	l2 = append(l2,"4")

	l3 := AllDifference(l1,l2)
	fmt.Println("l3:",l3)
	fmt.Println("l1:",l1)
	fmt.Println("l2:",l2)
	if len(l3) != 4 {
		t.Error("common.AllDifference() error,len: ",len(l3),"")
	}
}

func TestFindDirFiles(t *testing.T)  {
	fileList := FindDirFiles("../test/","../test/","a")
	l := len(fileList)
	fmt.Println(l)
	if l != 2 {
		t.Error("fail")
	} else {
		t.Log("success")
	}

}