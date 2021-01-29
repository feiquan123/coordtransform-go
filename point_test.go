package coordtransform

import "testing"

func TestPointRound(t *testing.T) {
	e := "114.000000000000,31.000000000000"
	r := p.Round(0).String()
	if r != e {
		t.Errorf("get:%s exp:%s\n", r, e)
	}
	e = "114.363000000000,30.517000000000"
	r = p.Round(3).String()
	if r != e {
		t.Errorf("get:%s exp:%s\n", r, e)
	}
}

func TestPointRoundString(t *testing.T) {
	e := "114.363,30.517"
	r := p.RoundString(3)
	if r != e {
		t.Errorf("get:%s exp:%s\n", r, e)
	}
	e = "114.3630890,30.5174110"
	r = p.RoundString(7)
	if r != e {
		t.Errorf("get:%s exp:%s\n", r, e)
	}
}
