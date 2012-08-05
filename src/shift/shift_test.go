package shift
import  (
	"fmt"
	"testing"
)

func TestShift( t *Testing.T )
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




