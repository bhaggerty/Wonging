package wonging

var (
	HiLo = map[string]int8{
		"2":  1,
		"3":  1,
		"4":  1,
		"5":  1,
		"6":  1,
		"7":  0,
		"8":  0,
		"9":  0,
		"10": -1,
		"J":  -1,
		"Q":  -1,
		"K":  -1,
		"A":  -1,
	}
	HiOpt1 = map[string]int8{
		"2":  0,
		"3":  1,
		"4":  1,
		"5":  1,
		"6":  1,
		"7":  0,
		"8":  0,
		"9":  0,
		"10": -1,
		"J":  -1,
		"Q":  -1,
		"K":  -1,
		"A":  0,
	}
	HiOpt2 = map[string]int8{
		"2":  1,
		"3":  1,
		"4":  2,
		"5":  2,
		"6":  1,
		"7":  1,
		"8":  0,
		"9":  0,
		"10": -2,
		"J":  -2,
		"Q":  -2,
		"K":  -2,
		"A":  0,
	}
	KO = map[string]int8{
		"2":  1,
		"3":  1,
		"4":  1,
		"5":  1,
		"6":  1,
		"7":  1,
		"8":  0,
		"9":  0,
		"10": -1,
		"J":  -1,
		"Q":  -1,
		"K":  -1,
		"A":  -1,
	}
	Omega2 = map[string]int8{
		"2":  1,
		"3":  1,
		"4":  2,
		"5":  2,
		"6":  2,
		"7":  1,
		"8":  0,
		"9":  -1,
		"10": -2,
		"J":  -2,
		"Q":  -2,
		"K":  -2,
		"A":  0,
	}
	Red7 = map[string]float32{
		"2":  1,
		"3":  1,
		"4":  1,
		"5":  1,
		"6":  1,
		"7":  0.5,
		"8":  0,
		"9":  0,
		"10": -1,
		"J":  -1,
		"Q":  -1,
		"K":  -1,
		"A":  -1,
	}
	ZenCount = map[string]int8{
		"2":  1,
		"3":  1,
		"4":  2,
		"5":  2,
		"6":  2,
		"7":  1,
		"8":  0,
		"9":  0,
		"10": -2,
		"J":  -2,
		"Q":  -2,
		"K":  -2,
		"A":  -1,
	}
)

type Counter struct {
	HiLo, HiOpt1, HiOpt2, KO, Omega2, ZenCount int8
	Red7                                       float32
}

func (c Counter) count(cardValue string) Counter {
	c.HiLo += HiLo[cardValue]
	c.HiOpt1 += HiOpt1[cardValue]
	c.HiOpt2 += HiOpt2[cardValue]
	c.KO += KO[cardValue]
	c.Omega2 += Omega2[cardValue]
	c.Red7 += Red7[cardValue]
	c.ZenCount += ZenCount[cardValue]
	return c
}
