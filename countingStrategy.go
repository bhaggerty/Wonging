package wonging

type CountingStrategy func(*Counter) float32

// strategies
func getHiLo(c *Counter) float32 {
	return float32(c.HiLo)
}

func getHiOpt1(c *Counter) float32 {
	return float32(c.HiOpt1)
}

func getHiOpt2(c *Counter) float32 {
	return float32(c.HiOpt2)
}

func getKO(c *Counter) float32 {
	return float32(c.KO)
}

func getOmega2(c *Counter) float32 {
	return float32(c.Omega2)
}

func getRed7(c *Counter) float32 {
	return float32(c.Red7)
}

func getZenCount(c *Counter) float32 {
	return float32(c.ZenCount)
}

func getAceFiveCount(c *Counter) float32 {
	return float32(c.AceFiveCount)
}

// strategies assignments
func randomCountingStrategy() (CountingStrategy, string) {
	strategies := []CountingStrategy{getHiLo, getHiOpt1, getHiOpt2, getKO, getOmega2, getRed7, getZenCount, getAceFiveCount}
	description := []string{"Hi Lo", "Hi Opt 1", "Hi Opt 2", "KO", "Omega 2", "Red 7", "Zen Count", "Ace/Five Count"}
	randomInt := randInt(0, len(strategies))
	return strategies[randomInt], description[randomInt]
}
