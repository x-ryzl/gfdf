package test

import (
	"gfdf/common"
	"testing"
)

func Benchmark_alldiffence(b *testing.B)  {
	var l1 = make([]string,b.N)
	var l2 = make([]string,b.N)
	for i:=1;i<b.N;i++ {
		l1[i] = "a"
	}
	for j:=1;j<b.N;j++ {
		l2[j] = "b"
	}
	common.AllDifference(l1,l2)
}

func Benchmark_diffence(b *testing.B)  {
	var l1 = make([]string,b.N)
	var l2 = make([]string,b.N)
	for i:=1;i<b.N;i++ {
		l1[i] = "a"
	}
	for j:=1;j<b.N;j++ {
		l2[j] = "b"
	}
	common.Difference(l1,l2)
}


func Benchmark_union(b *testing.B)  {
	var l1 = make([]string,b.N)
	var l2 = make([]string,b.N)
	for i:=1;i<b.N;i++ {
		l1[i] = "a"
	}
	for j:=1;j<b.N;j++ {
		l2[j] = "b"
	}
	common.Union(l1,l2)
}


func Benchmark_intersect(b *testing.B)  {
	var l1 = make([]string,b.N)
	var l2 = make([]string,b.N)
	for i:=1;i<b.N;i++ {
		l1[i] = "a"
	}
	for j:=1;j<b.N;j++ {
		l2[j] = "b"
	}
	common.Intersect(l1,l2)
}
