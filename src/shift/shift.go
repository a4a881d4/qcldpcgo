package qcldpc

import "fmt"

type shift struct {
	buf []uint32
}

var	length uint
var size uint
var mask uint32

func Init( length uint )	{
	length = length
	size = (length+31)/32
	left = size*32-length
	mask = shiftMy((0xffffffff),left)
}

func shiftMy( a uint32, s int32 ) uint32 {
	if s>=1	{
		a >>= 1
		a &= 0x7fffffff
		a >>= (s-1)
	}
	return a
}

func ( a *shift ) clean {
	for i:=0;i<size;i++ {
		a.buf[i] = 0
	}
}

func clean( a []shift )	{
	for i:=0;i<len(a);i++ {
		a[i].clean()
	}
}

func ( my *shift ) dump( a shift )	{
	for i:=0;i<size;i++ {
		my.buf[i]=a.buf[i];
	}
}

func cpy( from []shift, to []shift )	{
	for i:=0;i<to.length;i++ {
		to[i].dump(from[i])
	}
}

func (a *shift) weight() int	{
	i=0;
	for j:=0;j<length;j++ {
		if ((buf[j/32]>>(j%32))&1) == 1 {
			i++
		}
	}
	return i
}

func ( a *shift ) mac( b shift, s int ) {
	b = b.shiftH2L(s)
	for i:=0;i<size;i++ {
		a.buf[i]^=b.buf[i]
	}
}

func (a *shift) setXk( k uint )	{
		a.zero()
		a.buf[k/32] = 1<<(k%32);
}

func ( a *shift ) zero()	{
	for i:=0;i<size;i++ {
			a.buf[i]=0;
	}
}
	
func zero( a []shift )	{
	for i:=0;i<len(a);i++ {
		a[i].zero()
	}
}

func alloc( k uint ) []shift {
	var ret = make( []shift, k )
	for i:=0;i<k;i++ {
		ret[i]= new( shift )
	}
	zero(ret)
	return ret
}

func allocArray( a [][]int ) [][]shift {
	var ret = make( [][]shift, len(a) )
	for i:=0;i<len(a);i++ {
		ret[i]=alloc(a[i].length)
		for j:=0;j<len(a[i]);j++ {
			if a[i][j]!=length {
				ret[i][j].setXk((length-a[i][j])%length)
			}
		}
	}
	return ret
}

func reverse(shift a)	shift {
	b := new( shift )
	b.zero()
	for i=0;i<length;i++	{
		j := (length-i)%length
		c := (a.buf[j/32]>>(j%32))&1
		b.buf[i/32]|=c<<(i%32)
	}
	return b
}

func (a *shift) isXk() int {
	for i:=0;i<length;i++ {
		if ((buf[i/32]>>(i%32))&1)==1 {
			break
		}
	}
	if i==length {
		return -2
	}
	for j:=i+1;j<length;j++ {
		if ((buf[j/32]>>(j%32))&1)==1 {
			fmt.Printf("%d is one but %d is not zero, \"buf[%d]=%08x, buf[%d]=%08x\"\n",
					i,j,i/32,buf[i/32],j/32,buf[j/32])
			return -1
		}
	}
	return i
}
	
func mid( uint []a, begin int, end int, to int )	uint	{
	var r uint;
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
		r = a[posE]<<(31-leftE)
		r = shiftMy( r , (31-leftE+leftB))
		r = r<<(to-(end-begin))
	} else {
		r = shiftMy(a[posB],leftB)
		r |= a[posE]<<(to-leftE)
	}
	return r
}

func ( b *shift )shiftOneInt( a shift, s int )	{
	for i:=0;i<size-1;i++ {
		b.buf[i]=mid(a.buf,(i*32+s)%length,(i*32+31+s)%length,31)
	}
	b.buf[size-1]=mid(a.buf,((size-1)*32+s)%length,(length-1+s)%length,(length-1)%32)
	b.buf[size-1]&=mask
}

func ( a *shift ) shiftOne( s int ) shift{
	b := new( shift )
	for i:=0;i<size-1;i++ {
		b.buf[i]=mid(a.buf,(i*32+s)%length,(i*32+31+s)%length,31);
	}
	b.buf[size-1]=mid(a.buf,((size-1)*32+s)%length,(length-1+s)%length,(length-1)%32)
	b.buf[size-1]&=mask
	return b
}

func ( a *shift ) shiftH2L( s int ) shift	{
	return a.shiftOne(s)
}
	
func (a *shift) shiftL2H( s int ) shift {
		return a.shiftOne(length-s);
}
	
func (a *shift) shiftOneByExt( s int ) shift {
	r := new( shift )
	var temp = make([]int,length)
	for i:=0;i<length;i++ {
		temp[i] = (a.buf[i/32]>>(i%32))&1
	}
	r.zero()
	for i:=0;i<length;i++	{
		if temp[(i+s)%length]== 1 {
			r.buf[i/32]|=1<<(i%32)
		}
	}	
	return r
}
	
func ( a *shift ) Equ( b shift ) bool	{
	for i:=0;i<size;i++ {
		if a.buf[i]!=b.buf[i]	{
			return false
		}
	}
	return true
}

func ( a *shift ) print {
	c := 0;
	for i:=0;i<length;i++ {
		c<<=1
		c|=(buf[i/32]>>(i%32))&1
		if (i%4)==3	{
			fmt.Printf("%01x",c&0xf);
		}
	}
}

func ( a *shift ) printH2L	{
	c := 0;
	for i:=size-1;i>=0;i-- {
		fmt.Printf("%08x",a.buf[i])
	}
}

func printArray(a []shift, name string, comma int )	{
	if (comma&2) == 2 && name != null {
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

