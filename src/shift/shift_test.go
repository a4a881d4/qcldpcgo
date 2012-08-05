package shift

import (
    "testing"
    "math/rand"
)

func TestShiftOne(t *testing.T) {
	Init(104)
	var a = NewShift()
	a.SetXk(1)
	var b Shift
	b = *a
	for i:=0;i<104;i++ {
		b = b.ShiftOne(1)
		t.Logf("Shift One Test %d=%d\n",i,b.IsXk())
		if b.IsXk()!=(1-1-i+104)%104 {
			t.Fail()
		}
	}
}

func TestShiftRandom(t *testing.T) {
	length := 104
	Init(length)
	var a = NewShift()
	for i:=0;i<104;i++ {
		pos := rand.Intn(length)
		s := rand.Intn(length)
		a.SetXk(pos)	
		b := a.ShiftOne(s)
		t.Logf("Shift random Test [%d]: %d -> %d=%d\n",i,pos,s,b.IsXk())
		if b.IsXk()!=(pos-s+length)%length {
			t.Fail()
		}
	}
}

func TestShiftClean(t *testing.T) {
	length := 104
	Init(length)
	var a = NewShift()
	a.SetXk(1)
	if a.IsXk()!=1 {
		t.Fail()
	}
	a.clean();
	if a.IsXk()!=-2 {
		t.Fail()
	}
}
 