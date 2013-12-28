package wonging

import (
	"testing"
)

func Test_Count(t *testing.T) {
	counter := new(Counter).initialize()
	counter.count("2").count("7").count("10").count("K").count("A")
	if counter.AceFiveCount != -1 {
		t.Error("Count() for [AceFiveCount] did not work as expected.")
	} else {
		t.Log("Count() for [AceFiveCount] test passed")
	}

	if counter.HiLo != -2 {
		t.Error("Count() for [HiLo] did not work as expected.")
	} else {
		t.Log("Count() for [HiLo] test passed")
	}

	if counter.HiOpt1 != -2 {
		t.Error("Count() for [HiOpt1] did not work as expected.")
	} else {
		t.Log("Count() for [HiOpt1] test passed")
	}

	if counter.HiOpt2 != -2 {
		t.Error("Count() for [HiOpt2] did not work as expected.")
	} else {
		t.Log("Count() for [HiOpt2] test passed")
	}

	if counter.KO != -1 {
		t.Error("Count() for [KO] did not work as expected.")
	} else {
		t.Log("Count() for [KO] test passed")
	}

	if counter.Omega2 != -2 {
		t.Error("Count() for [Omega2] did not work as expected.")
	} else {
		t.Log("Count() for [Omega2] test passed")
	}

	if counter.Red7 != -1.5 {
		t.Error("Count() for [Red7] did not work as expected.")
	} else {
		t.Log("Count() for [Red7] test passed")
	}

	if counter.ZenCount != -3 {
		t.Error("Count() for [ZenCount] did not work as expected.")
	} else {
		t.Log("Count() for [ZenCount] test passed")
	}
}
