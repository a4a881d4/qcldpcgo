package qcldpc

type shift struct {
	buf []uint32;
	length uint;
	size uint;
	mask uint32;
}

func (a shfit)init	{
	a.size := (length+31)/32
	left := size*32-length
	a.mask := shiftMy((0xffffffff),left)
}

func shiftMy( a uint32, s int 32 ) uint32 {
	if s>=1	{
		a := a>>1;
		a &= 0x7fffffff;
		a := a>>(s-1);
	}
	return a;
}


