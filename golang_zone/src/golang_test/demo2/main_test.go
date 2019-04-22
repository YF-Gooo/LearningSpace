package main
import (
	"testing"
	"fmt"
)

func TestMain(m *testing.M){
	fmt.Println("test main first")
	m.Run()
}

func TestPrint(t *testing.T){
	t.Run("a1",func(t *testing.T){fmt.Println("a1")})
	t.Run("a2",func(t *testing.T){fmt.Println("a2")})
	t.Run("a3",func(t *testing.T){fmt.Println("a3")})
}

func testPrint1to20(t *testing.T){
	res:=Print1to20()
	fmt.Println("hey")
	if res!=210{
		t.Errorf("wrong result of Print1to20")
	}
}

func TestPrint2(t *testing.T){
	t.Run("testPrint1to20",testPrint1to20)
}

func BenchmarkAll(b *testing.B){
	for n:=0;n<b.N;n++{
		Print1to20()
	}
}