package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func init(){
	PreviousSolution(ioutil.Discard)
	NewSolution(ioutil.Discard)
}

func TestSolution(t *testing.T)  {
	prevOut:=new(bytes.Buffer)
	PreviousSolution(prevOut)
	prevResult:=prevOut.String()

	newOut:=new(bytes.Buffer)
	NewSolution(newOut)
	newResult:=newOut.String()

	if prevResult != newResult {
		t.Errorf("results not match\nGot:\n%v\nExpected:\n%v", newResult, prevResult)
	}
}

func BenchmarkPreviousSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
     PreviousSolution (ioutil.Discard)
	}
}

func BenchmarkNewSolution(b *testing.B) {
	for i:=0; i<b.N; i++ {
		NewSolution(ioutil.Discard)
	}
}

//go test -bench . -benchmem -cpuprofile=cpu.out -memprofile=mem.out -memprofilerate=1
//go tool pprof testquest.test.exe mem.out
//go tool pprof testquest.test.exe cpu.out

