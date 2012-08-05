package main 

import (
	"pkg/shift"
	"fmt"
)

func main() {

	shift.Init(104)
	var a = shift.NewShift()
	a.SetXk(100)
	var b shift.Shift
	b = *a
	for i:=0;i<104;i++ {
		b = b.ShiftOne(1)
		fmt.Printf("%d\n",b.IsXk())
	}
}

