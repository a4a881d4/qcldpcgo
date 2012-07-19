package qcldpc

import 'fmt';

type shift struct {
	buf []uint32;
}

var	length uint;
var size uint;
var mask uint32;

func Init( length uint )	{
	size := (length+31)/32
	left := size*32-length
	mask := shiftMy((0xffffffff),left)
}

func shiftMy( a uint32, s int 32 ) uint32 {
	if s>=1	{
		a := a>>1;
		a &= 0x7fffffff;
		a := a>>(s-1);
	}
	return a;
}

func (a shift) clean {
	for i:=0;i<size;i++ {
		a.buf[i] := 0
	}
}

func clean( a []shift )	{
	for i:=0;i<len(a);i++ {
		a[i].clean()
	}
}

func ( my shift ) dump( a shift )	{
	for i:=0;i<size;i++ {
		my.buf[i]=a.buf[i];
	}
}

func cpy(Shift[] from, Shift[] to )	{
	for i:=0;i<to.length;i++ {
		to[i].dump(from[i])
	}
}
func (a shift) weight() int	{
	i=0;
	for j:=0;j<length;j++ {
		if ((buf[j/32]>>(j%32))&1) == 1 {
			i++;
		}
	}
	return i;
}

func ( a shift ) mac( a shift, s int ) {
	b := a.shiftH2L(s)
	for i:=0;i<size;i++ {
		a.buf[i]^=b.buf[i];
}

func (a shift) setXk( k uint )	{
		a.zero();
		a.buf[k/32] := 1<<(k%32);
}

func ( a shift ) zero()	{
	for i:=0;i<size;i++ {
			a.buf[i]:=0;
	}
}
	
func zero( a []shift )	{
	for i:=0;i<len(a);i++ {
		a[i].zero()
	}
}

func alloc( k uint ) []shift {
	ret := new [k]shift;
	for i:=0;i<k;i++ {
		ret[i]= make( shift );
	}
	zero(ret);
	return ret;
}

func alloc( a [][]int) [][]shift {
	int i,j;
	ret := new [len(a)][]shift;
	for i:=0;i<len(a);i++ {
		ret[i]=Shift.alloc(a[i].length)
		for j=0;j<len(a[i]);j++ {
			if a[i][j]!=length {
					ret[i][j].setXk((length-a[i][j])%length);
			}
		}
	}
	return ret;
}

func reverse(shift a)	shift{
	b := new shift;
	b.zero();
	for i=0;i<length;i++	{
		j := (length-i)%length;
		c := (a.buf[j/32]>>(j%32))&1;
		b.buf[i/32]|=c<<(i%32);
	}
	return b;
}

func (a shift) isXk() int {
	for i:=0;i<length;i++ {
		if ((buf[i/32]>>(i%32))&1)==1 break;
	}
	if i==length return -2;
	for j:=i+1;j<length;j++ {
		if ((buf[j/32]>>(j%32))&1)==1 {
			fmt.printf("%d is one but %d is not zero, \"buf[%d]=%08x, buf[%d]=%08x\"\n",
					i,j,i/32,buf[i/32],j/32,buf[j/32]);
			return -1;
		}
	}
	return i;
}
	
func mid( uint []a, int begin, int end, int to )	uint	{
	var r uint = 0;
	if begin>end {
		r := mid(a,0,end,to);
		r |= mid(a,begin,length-1,length-1-begin);
		return r;
	}
	posE := end/32;
	leftE := end%32;
	posB := begin/32;
	leftB := begin%32;
	if posE==posB	{
		r := a[posE]<<(31-leftE);
		r := shiftMy( r , (31-leftE+leftB));
		r := r<<(to-(end-begin));
	}
	else {
		r := shiftMy(a[posB],leftB);
		r |= a[posE]<<(to-leftE);
	}
	return r;
}

func ( b shift )shiftOneInt( a shift, s int )	{
	for i:=0;i<size-1;i++ {
		b.buf[i]=mid(a.buf,(i*32+s)%length,(i*32+31+s)%length,31);
	}
	b.buf[size-1]=mid(a.buf,((size-1)*32+s)%length,(length-1+s)%length,(length-1)%32);
	b.buf[size-1]&=mask;
}

func ( a shift ) shiftOne( int s ) shift{
	b := new shift;
	for i:=0;i<size-1;i++ {
		b.buf[i]=mid(a.buf,(i*32+s)%length,(i*32+31+s)%length,31);
	}
	b.buf[size-1]=mid(a.buf,((size-1)*32+s)%length,(length-1+s)%length,(length-1)%32);
	b.buf[size-1]&=mask;
	return b;
}

func ( a shift ) shiftH2L( s int ) shift	{
	return a.shiftOne(s);
}
	
func (a shift) shiftL2H( s int ) shift {
		return a.shiftOne(length-s);
}
	
func (a shift) shiftOneByExt( s int ) shift {
	r := make(shift);
	temp := new [length]int;
	for i:=0;i<length;i++	temp[i] := (a.buf[i/32]>>(i%32))&1;
	for i:=0;i<size;i++ r.buf[i]:=0;
	for i:=0;i<length;i++	{
		if temp[(i+s)%length]== 1 r.buf[i/32]|=1<<(i%32);
	}	
	return r;
}
	
func	( a shift ) Equ( b shift ) bool	{
	for i:=0;i<size;i++ {
		if a.buf[i]!=b.buf[i]	{
			return false;
		}
	}
	return true;
}

func ( a shift ) print {
	c := 0;
	for i:=0;i<length;i++ {
		c<<=1;
		c|=(buf[i/32]>>(i%32))&1;
		if (i%4)==3	{
			fmt.printf("%01x",c&0xf);
		}
	}
}

func ( a shift ) printH2L	{
	c := 0;
	for i:=size-1;i>=0;i-- {
		fmt.printf("%08x",a.buf[i]);
	}
}

func printArray(Shift[] a, String name, int comma )	{
	if (comma&2) == 2 && name != null {
		fmt.println(name);
	}
	for i:=0;i<len(a);i++ {
		a[i].print();
		if (comma&1) == 1 {
				fmt.printf(",");
		}
	}
	if (comma&2) == 2 {
			fmt.printf("\n");
	}
}

	Shift TestShift( int s )
	{
		int j;
		Shift c = this.shiftOne(s);
		Shift b = this.shiftOneByExt(s);
		if( !b.Equ(c) )
		{
			System.out.printf("b = ");
			for( j=0;j<Shift.size;j++ )
				System.out.printf("%08x ",b.buf[j]);
			System.out.printf("\n");
			System.out.printf("c = ");
			for( j=0;j<Shift.size;j++ )
				System.out.printf("%08x ",c.buf[j]);
			System.out.printf("\n");
		}
		
		return b;
	}
	static int TestXk(int k, int s)
	{
		Shift a = new Shift();
		a.setXk(k);
		a = a.TestShift(s);
		return a.isXk();
	}
	private void setBit(int k, int pos)
	{
		
		buf[pos/32] |= (k<<(pos%32));
	}
	static Shift[] msg2shift(byte[] msg)
	{
		int len = msg.length*8/length;
		Shift[] ret = new Shift[len];
		for( int i=0;i<len; i++ )
			ret[i] = new Shift();
		Shift.msg2shift(ret, msg, len*length);
		return ret;
	}
	static int msg2shift( Shift[] info, byte[] msg, int k )
	{
		if( Shift.length*info.length != k )
			return -1;
		int i;
		for( i=0;i<info.length;i++ )
			info[i].clean();
		for( i=0;i<k;i++ )
			info[i/Shift.length].setBit(((int)msg[i/8]>>(i%8))&1,i%Shift.length);
		return Shift.length*info.length;
	}
	public static int shift2msg( Shift[] info, byte[] msg, int k )
	{
		int i;
		for( i=0;i<info.length*Shift.length/8;i++ )
			msg[k+i]=0;
		for( i=0;i<info.length*Shift.length;i++ )
		{
			int pos = i%Shift.length;
			int bit = info[i/Shift.length].buf[pos/32]>>(pos%32);
			bit &= 1;
			msg[k+i/8]|=(bit<<i%8);
		}
		return info.length*Shift.length;
	}
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		int i,j,k;
		Shift.length = 136;
		Shift.init();
		
		Shift a = new Shift();
		
		for( i=0;i<10000;i++ )
		{
			
			for( j=0;j<Shift.size;j++ )
			{
				a.buf[j] = (int) Math.round(Math.random()*2147483647.);
			}
			a.buf[Shift.size-1] &= Shift.mask;
			
			a.TestShift(i%136);
		
			if( i%1000 == 0 )
				System.out.printf("*");
		}
		System.out.printf("\n");
		
		a.setXk(14);
		k=2;
		for( i=0;i<10000;i++ )
		{
			k = Shift.TestXk(k,i%136);
		}
		
		System.out.printf("\n");
		a.print();
		System.out.printf("\n");
		
		
	}
}




