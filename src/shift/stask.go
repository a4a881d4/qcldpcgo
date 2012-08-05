package shift

import (
	"fmt"
)


struct stMac {
	in int
	mac int
	ro int
}

type build int32
struct STask {
	const build (
	    _ = iota 
		matrix
		matrixT
		guass
		density
	)
	task []stMac
}

func NewSTask( rs [][]int, row int, col int, bl int, t build  ) *STask {
	switch( t ) {
			case matrix:
				return matrixBuild( rs, row, col, bl, t )
			case matrixT:
				return matrixBuild( rs, row, col, bl, t )
			case guass:
				return guassBuild( rs, row, col, bl )
			default:
				taskLen = 0
	}
}

func DensitySTask( rs [][]*ShiftB, int row, int col ) *STask {
	return densityBuild( rs, row, col );
}
	
func matrixBuild( c [][]int, row int, col int, bl int, t build  ) *STask {
	aT := new( STask )
	aT.taskLen = sumWeight( c, row, col, bl, t )
	task = make( []stMac, taskLen )
	pos := 0
	for i:=0;i<row;i++ {
		for( j:=0;j<col;j++ ) {
			if( c[i][j]!=bl ) {
				aT.task[pos].ro=c[i][j]
				if( t==build.matrix ) {
					aT.task[pos].in = j
					aT.task[pos].mac = i
				} else {
					aT.task[pos].in = i
					aT.task[pos].mac = j
				}
				pos++
			}
		}
	}
	return aT
}

func guassBuild( c [][]int, row int, col int, bl int  ) *STask {
	aT := new( STask )
	pos:=0
	aT.taskLen = sumWeight( c, row, col, bl, build.guass );
	task = make( []stMac, taskLen )
	for i:=0;i<col;i++ {
		for j:=i+1;j<row;j++ {
			if( c[i][j]!=bl ) {
				aT.task[pos].ro=c[i][j]
				aT.task[pos].in = i
				aT.task[pos].mac = j
				pos++
			}
		}
	}
	return aT
}

func densityBuild( c [][]*ShiftB, row int, col int ) *STask {
	aT := new( STask )
	pos:=0
	aT.taskLen = densitySumWeight( c, row, col )
	task = make( []stMac, taskLen )
	for i:=0;i<row;i++ {
		for j:=0;j<col;j++ 	{
			s := c[i][j]
			for l=0;l<length;l++ {
				if( s.b.Bit(l)==1 {
					aT.task[pos].in=i
					aT.task[pos].mac=j
					aT.task[pos].ro=l
					aT.pos++
				}
			}
		}
	}
	return aT
}
	
func densitySumWeight( c [][]*ShiftB, row int, col int ) int {
	count := 0;
	for i:=0;i<row;i++ {
		for j:=0;j<col;j++ {
			count += c[i][j].weight()
		}
	}	
	return count
}
	
func sumWeight( c [][] int, row int, col int, bl int, t build ) int {
	count:=0
	switch( t )	{
		case matrix,matrixT:
			for i:=0;i<row;i++ {
				for j:=0;j<col;j++ {
					if c[i][j]!=bl {
						count++
					}
				}
			}	
		case guass:
			for i:=0;i<col;i++ {
				for j:=i+1;j<row;j++ ) {
					if c[i][j]!=bl {
						count++
					}
				}
			}
		default:
	}
	return count
}

	void  doit( Shift in[], Shift out[] )
	{
		int i;
		for( i=0;i<taskLen;i++ )
		{
			out[task[i].mac].mac(in[task[i].in], task[i].ro);
		}
	}
	void  doit( Shift in[], Shift out[], int offsetIn, int offsetOut )
	{
		int i;
		for( i=0;i<taskLen;i++ )
		{
			out[task[i].mac+offsetOut].mac(in[task[i].in+offsetIn], task[i].ro);
		}
	}
	void printfTask()
	{
		System.out.printf("task length =%d\n",taskLen);
		for( int i=0;i<taskLen;i++ )
		{
			System.out.printf("task[%d] is shift[%d] %d mac[%d]\n",i,task[i].in,task[i].ro,task[i].mac);
		}
	}
	void shiftOneTest( int s)
	{
		System.out.printf("Test SHift Once\n");
		Shift[] testUin = Shift.alloc(1);
		Shift[] testUout = Shift.alloc(1);
		
		int[][] c={ {s} };
		matrixBuild(c,1,1,Shift.length,build.matrix);
		printfTask();
		for( int testP = 0; testP<Shift.length; testP+=13 )
		{
			testUin[0].setXk(testP);
			Shift.zero(testUout);
			System.out.printf("\n");
			Shift.printArray(testUin,"In :",2);
			doit(testUin,testUout);
			System.out.printf("test at %d shift %d to %d \n",testUin[0].isXk(),s,testUout[0].isXk());
			Shift.printArray(testUout,"Out :",2);
			if( (testUin[0].isXk()-testUout[0].isXk()+Shift.length)%Shift.length!=s )
				System.out.printf("shift error at %d \n",testP );
		}
	}
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Shift.length = 136;
		Shift.init();
		STask ttu = new STask();
		ttu.shiftOneTest(3);
	}

}