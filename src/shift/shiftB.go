package shift

import (
	"fmt"
	"math/big"
	"math/rand"
)

var rnd *rand.Rand

type ShiftB struct {
	b *big.Int
}

var maskB *big.Int

func NewShiftB() *ShiftB {
	a := new( ShiftB )
	a.b = big.NewInt(0)
	return a
}

func InitB( inLen int )	{
	length = inLen
	maskB = big.NewInt(1)
	maskB.Lsh(maskB,uint(length))
	maskB.Sub(maskB,big.NewInt(1))
	src := rand.NewSource(1)
	rnd = rand.New(src) 
}

func (a *ShiftB)shiftMy( s int ) {
	a.b.Rsh(a.b,uint(s))
}

func ( a *ShiftB ) clean() {
	a.b = big.NewInt(0)
}

func cleanB( a []*ShiftB ) {
	for i:=0;i<len(a);i++ {
		a[i].clean()
	}
}

func ( my *ShiftB ) dump( a *ShiftB )	{
	my.b.Set(a.b)
}

func cpyB( from []*ShiftB, to []*ShiftB )	{
	for i:=0;i<len(to);i++ {
		to[i].dump(from[i])
	}
}

func (a *ShiftB) weight() int	{
	i:=0
	for j:=0;j<int(length);j++ {
		if a.b.Bit(j) == 1 {
			i++
		}
	}
	return i
}

func ( a *ShiftB ) Mac( b *ShiftB, s int ) {
	b = b.ShiftH2L(s)
	a.b=a.b.Xor(a.b,b.b)
}

func (a *ShiftB) SetXk( k int )	{
	a.clean()
	a.b.SetBit(a.b,k,1)
}

func ( a *ShiftB ) zero()	{
	a.b= big.NewInt(0)
}
	
func ( c *ShiftB ) zeroArray( a []*ShiftB )	{
	for i:=0;i<len(a);i++ {
		a[i].zero()
	}
}

func ( c *ShiftB ) alloc( k int ) []*ShiftB {
	var ret = make( []*ShiftB, k )
	for i:=0;i<k;i++ {
		ret[i]=NewShiftB()
	}
	ret[0].zeroArray(ret)
	return ret
}

func ( c *ShiftB ) alloc2D( a [][]int ) [][]*ShiftB {
	temp := NewShiftB()
	var ret = make( [][]*ShiftB, len(a) )
	for i:=0;i<len(a);i++ {
		ret[i]=temp.alloc(len(a[i]))
		for j:=0;j<len(a[i]);j++ {
			if a[i][j]!=int(length) {
				ret[i][j].SetXk((length-int(a[i][j]))%length)
			}
		}
	}
	return ret
}

func ( a *ShiftB )reverse()	*ShiftB {
	b := NewShiftB()
	b.zero()
	var i int;
	for i=0;i<length;i++	{
		j := (length-i)%length
		c := a.b.Bit(j)
		b.b.SetBit(b.b,i,c)
	}
	return b
}

func (a *ShiftB) IsXk() int {
	var i int;
	for i=0;i<int(length);i++ {
		if a.b.Bit(i)==1 {
			break
		}
	}
	if i==int(length) {
		return -2
	}
	for j:=i+1;j<int(length);j++ {
		if a.b.Bit(j)==1 {
			return -1
		}
	}
	return i
}

func ( b *ShiftB )ShiftOneInt( a *ShiftB, s int )	{
	if s>0 {
		sh := length-s
		c := big.NewInt(0)
		c = c.Set(a.b)
//		fmt.Printf("c dump from a %x\n",c)
		c.Lsh(c,uint(sh))
//		fmt.Printf("c shift left %x %d\n",c,sh)
		c = c.And(c,maskB)
//		fmt.Printf("c mask with %x %x\n",maskB,c)
		d := big.NewInt(0)
		d = a.b.Set(a.b)
//		fmt.Printf("d dump from a %x\n",d)
		d.Rsh(d,uint(s))
//		fmt.Printf("d shift right %x %d\n",d,s)
		b.b = c.Or(c,d)
	}	else if s==0 {
		b.b = a.b.Set(a.b)
	}
}

func ( a *ShiftB ) ShiftOne( s int ) *ShiftB{
	b := NewShiftB()
	b.ShiftOneInt(a,s)
	return b
}

func ( a *ShiftB ) ShiftH2L( s int ) *ShiftB	{
	return a.ShiftOne(s)
}
	
func (a *ShiftB) ShiftL2H( s int ) *ShiftB {
		return a.ShiftOne(length-s);
}
	
func (a *ShiftB) ShiftOneByExt( s int ) *ShiftB {
	r := NewShiftB()
	for i:=0;i<length;i++ {
		c := a.b.Bit(i)
		r.b.SetBit(r.b,(length+i-s)%length,c)
	}
	return r
}
	
func ( a *ShiftB ) Equ( b *ShiftB ) int	{
	return a.b.Cmp(b.b)
}

func ( a *ShiftB ) print() {
	fmt.Printf("%x\n",a.b)
}

func printArrayB(a []*ShiftB, name string, comma int )	{
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

func ( a *ShiftB ) setB( b *Shift ) {
	a.clean()
	for j:=0;j<int(length);j++ {
		if ((b.buf[j/32]>>uint(j%32))&1) == 1 {
			a.b.SetBit(a.b,j,1)
		}
	}
}	

func ( a *ShiftB ) set( b *Shift ) {
	a.clean()
	for j:=0;j<int(length);j++ {
		if a.b.Bit(j)==1 {
			b.buf[j/32]|=1<<uint(j%32)
		}
	}
}	

func ( a *ShiftB ) random() {
	a.b = a.b.Rand(rnd,maskB)
}