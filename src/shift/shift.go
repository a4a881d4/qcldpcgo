package shift

import (
	"fmt"
)

type Shift struct {
	buf []uint32
}

var	length int
var size int
var mask uint32

func NewShift() *Shift{
	a := new( Shift )
	a.buf = make( []uint32, size )
	return a
}

func Init( inLen int )	{
	length = inLen
	size = (length+31)/32
	var left int
	left = int(size*32-length)
	mask = shiftMy((0xffffffff),left)
}

func shiftMy( a uint32, s int ) uint32 {
	if s>=1	{
		a >>= 1
		a &= 0x7fffffff
		a >>= uint(s-1)
	}
	return a
}

func ( a *Shift ) clean() {
	for i:=0;i<int(size);i++ {
		a.buf[i] = 0
	}
}

func clean( a []Shift ) {
	for i:=0;i<len(a);i++ {
		a[i].clean()
	}
}

func ( my *Shift ) dump( a Shift )	{
	for i:=0;i<int(size);i++ {
		my.buf[i]=a.buf[i];
	}
}

func cpy( from []Shift, to []Shift )	{
	for i:=0;i<len(to);i++ {
		to[i].dump(from[i])
	}
}

func (a *Shift) weight() int	{
	i:=0;
	for j:=0;j<int(length);j++ {
		if ((a.buf[j/32]>>uint(j%32))&1) == 1 {
			i++
		}
	}
	return i
}

func ( a *Shift ) Mac( b Shift, s int ) {
	b = b.ShiftH2L(s)
	for i:=0;i<int(size);i++ {
		a.buf[i]^=b.buf[i]
	}
}

func (a *Shift) SetXk( k int )	{
		a.zero()
		a.buf[k/32] = 1<<uint(k%32);
}

func ( a *Shift ) zero()	{
	for i:=0;i<int(size);i++ {
			a.buf[i]=0;
	}
}
	
func zero( a []Shift )	{
	for i:=0;i<len(a);i++ {
		a[i].zero()
	}
}

func alloc( k int ) []Shift {
	var ret = make( []Shift, k )
	zero(ret)
	return ret
}

func alloc2D( a [][]int ) [][]Shift {
	var ret = make( [][]Shift, len(a) )
	for i:=0;i<len(a);i++ {
		ret[i]=alloc(len(a[i]))
		for j:=0;j<len(a[i]);j++ {
			if a[i][j]!=int(length) {
				ret[i][j].SetXk((length-int(a[i][j]))%length)
			}
		}
	}
	return ret
}

func reverse( a Shift )	Shift {
	b := NewShift()
	b.zero()
	var i int;
	for i=0;i<length;i++	{
		j := (length-i)%length
		c := (a.buf[j/32]>>uint(j%32))&1
		b.buf[i/32]|=c<<uint(i%32)
	}
	return *b
}

func (a *Shift) IsXk() int {
	var i int;
	for i=0;i<int(length);i++ {
		if ((a.buf[i/32]>>uint(i%32))&1)==1 {
			break
		}
	}
	if i==int(length) {
		return -2
	}
	for j:=i+1;j<int(length);j++ {
		if ((a.buf[j/32]>>uint(j%32))&1)==1 {
			fmt.Printf("%d is one but %d is not zero, \"buf[%d]=%08x, buf[%d]=%08x\"\n",
					i,j,i/32,a.buf[i/32],j/32,a.buf[j/32])
			return -1
		}
	}
	return i
}
	
func mid( a []uint32, begin int, end int, to int )	uint32	{
	var r uint32;
	if begin>end {
		r = mid(a,0,end,to);
		r |= mid(a,begin,length-1,length-1-begin);
		return r
	}
	posE := end/32;
	leftE := end%32;
	posB := begin/32;
	leftB := begin%32;
	if posE==posB	{
		r = a[posE]<<uint(31-leftE)
		r = shiftMy( r , (31-leftE+leftB))
		r = r<<uint(to-(end-begin))
	} else {
		r = shiftMy(a[posB],leftB)
		r |= a[posE]<<uint(to-leftE)
	}
	return r
}

func ( b *Shift )ShiftOneInt( a Shift, s int )	{
	for i:=0;i<size-1;i++ {
		b.buf[i]=mid(a.buf,(i*32+s)%length,(i*32+31+s)%length,31)
	}
	b.buf[size-1]=mid(a.buf,((size-1)*32+s)%length,(length-1+s)%length,(length-1)%32)
	b.buf[size-1]&=mask
}

func ( a *Shift ) ShiftOne( s int ) Shift{
	b := NewShift()
	for i:=0;i<size-1;i++ {
		b.buf[i]=mid(a.buf,(i*32+s)%length,(i*32+31+s)%length,31);
	}
	b.buf[size-1]=mid(a.buf,((size-1)*32+s)%length,(length-1+s)%length,(length-1)%32)
	b.buf[size-1]&=mask
	return *b
}

func ( a *Shift ) ShiftH2L( s int ) Shift	{
	return a.ShiftOne(s)
}
	
func (a *Shift) ShiftL2H( s int ) Shift {
		return a.ShiftOne(length-s);
}
	
func (a *Shift) ShiftOneByExt( s int ) Shift {
	r := NewShift()
	var temp = make([]uint32,length)
	for i:=0;i<length;i++ {
		temp[i] = (a.buf[i/32]>>uint(i%32))&1
	}
	r.zero()
	for i:=0;i<length;i++	{
		if temp[(i+s)%length]==1 {
			r.buf[i/32]|=1<<uint(i%32)
		}
	}	
	return *r
}
	
func ( a *Shift ) Equ( b Shift ) bool	{
	for i:=0;i<size;i++ {
		if a.buf[i]!=b.buf[i]	{
			return false
		}
	}
	return true
}

func ( a *Shift ) print() {
	c := uint32(0);
	for i:=0;i<length;i++ {
		c<<=1
		c|=(a.buf[i/32]>>uint(i%32))&1
		if (i%4)==3	{
			fmt.Printf("%01x",c&0xf);
		}
	}
}

func ( a *Shift ) printH2L()	{
	for i:=size-1;i>=0;i-- {
		fmt.Printf("%08x",a.buf[i])
	}
}

func printArray(a []Shift, name string, comma int )	{
	if (comma&2) == 2 && name != "" {
		fmt.Println(name);
	}
	for i:=0;i<len(a);i++ {
		a[i].print()
		if (comma&1) == 1 {
				fmt.Printf(",");
		}
	}
	if (comma&2) == 2 {
			fmt.Printf("\n");
	}
}

