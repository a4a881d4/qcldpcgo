package shift

import (
    "testing"
    "math/rand"
)

func TestShiftBOne(t *testing.T) {
	length := 104
	InitB(length)
	var a = NewShiftB()
	a.SetXk(1)
	var b *ShiftB
	b = a
	for i:=0;i<length;i++ {
		b.print()
		b = b.ShiftOne(1)
		t.Logf("Shift One Test %d = %d\n",i,b.IsXk())
		if b.IsXk()!=(1-1-i+length)%length {
			b.print()
			t.Fail()
		}
	}
}

func TestShiftBRandom(t *testing.T) {
	length := 104
	InitB(length)
	var a = NewShiftB()
	for i:=0;i<length;i++ {
		pos := rand.Intn(length)
		s := rand.Intn(length)
		a.SetXk(pos)	
		b := a.ShiftOne(s)
		t.Logf("Shift random Test [%d]: %d -> %d = %d\n",i,pos,s,b.IsXk())
		if b.IsXk()!=(pos-s+length)%length {
			a.print()
			b.print()
			t.Fail()
		}
	}
}

func TestShiftBClean(t *testing.T) {
	length := 104
	InitB(length)
	var a = NewShiftB()
	a.SetXk(1)
	if a.IsXk()!=1 {
		t.Fail()
	}
	a.clean()
	if a.IsXk()!=-2 {
		t.Fail()
	}
}

func TestShiftTwo( t *testing.T ) {
	length := 104
	InitB(length)
	Init(length)
	for i:=0;i<length;i++ {
		s := rand.Intn(length)
		var a = NewShiftB()
		a.random()
		Sa := NewShift()
		a.set(Sa)
		Sb := Sa.ShiftOne(s)
		a = a.ShiftOne(s)
		var b = NewShiftB()
		b.setB(&Sb)
		if a.Equ(b)!= 0 {
			a.print()
			Sb.print()
			t.Fail()
		}
	}
}

func TestShiftTwoMac( t *testing.T ) {
	length := 104
	InitB(length)
	Init(length)
	for i:=0;i<length;i++ {
		s := rand.Intn(length)
		var a = NewShiftB()
		var b = NewShiftB()
		a.random()
		b.random()
		Sa := NewShift()
		a.set(Sa)
		Sb := NewShift()
		b.set(Sb)
		Sa.Mac(*Sb,s)
		a.Mac(b,s)
		var c = NewShiftB()
		c.setB(Sa)
		if a.Equ(c)!= 0 {
			a.print()
			Sa.print()
			t.Fail()
		}
	}
}
  