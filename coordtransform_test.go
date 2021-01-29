package coordtransform

import "testing"

var (
	l     = 7
	p     = Point{114.363089, 30.517411}
	pBDMC = Point{12730979.361410993, 3548336.555792139}
	pMC   = Point{12730840.83125368, 3569413.5156310988}

	expBD09toGCJ02  = Point{114.356659806424, 30.511068245460944}
	expGCJ02toBD09  = Point{114.36951652929112, 30.523754384540837}
	expWGS84toGCJ02 = Point{114.36861383019243, 30.515078352018758}
	expGCJ02toWGS84 = Point{114.35756416980757, 30.51974364798124}
	expBD09toWGS84  = Point{114.35114337613005, 30.513412637358606}
	expWGS84toBD09  = Point{114.3750500836037, 30.521392123098003}
	expBD09toBD09MC = Point{12730979.354966495, 3548336.4718212215}
	expDB09MctoBD09 = Point{114.36308905789122, 30.517411569328928}
	expToMercator   = Point{12730840.83125368, 3570233.0974655827}
	expFromMercator = Point{114.363089, 30.511068245460944}
)

func TestBD09toGCJ02(t *testing.T) {
	tp := BD09toGCJ02(p).Round(l)
	if !expBD09toGCJ02.Round(l).Equals(tp) {
		t.Errorf("BD09toGCJ02 err, get:%v exp:%v \n", expBD09toGCJ02.Round(l), tp)
	}
}

func TestGCJ02toBD09(t *testing.T) {
	tp := GCJ02toBD09(p).Round(l)
	if !expGCJ02toBD09.Round(l).Equals(tp) {
		t.Errorf("GCJ02toBD09 err, get:%v exp:%v \n", expGCJ02toBD09.Round(l), tp)
	}
}

func TestWGS84toGCJ02(t *testing.T) {
	tp := WGS84toGCJ02(p).Round(l)
	if !expWGS84toGCJ02.Round(l).Equals(tp) {
		t.Errorf("WGS84toGCJ02 err, get:%v exp:%v \n", expWGS84toGCJ02.Round(l), tp)
	}
}

func TestGCJ02toWGS84(t *testing.T) {
	tp := GCJ02toWGS84(p).Round(l)
	if !expGCJ02toWGS84.Round(l).Equals(tp) {
		t.Errorf("GCJ02toWGS84 err, get:%v exp:%v \n", expGCJ02toWGS84.Round(l), tp)
	}
}

func TestBD09toWGS84(t *testing.T) {
	tp := BD09toWGS84(p).Round(l)
	if !expBD09toWGS84.Round(l).Equals(tp) {
		t.Errorf("BD09toWGS84 err, get:%v exp:%v \n", expBD09toWGS84.Round(l), tp)
	}
}
func TestWGS84toBD09(t *testing.T) {
	tp := WGS84toBD09(p).Round(l)
	if !expWGS84toBD09.Round(l).Equals(tp) {
		t.Errorf("WGS84toBD09 err, get:%v exp:%v \n", expWGS84toBD09.Round(l), tp)
	}
}
func TestBD09toBD09MC(t *testing.T) {
	tp := BD09toBD09MC(p).Round(l)
	if !expBD09toBD09MC.Round(l).Equals(tp) {
		t.Errorf("BD09toBD09MC err, get:%v exp:%v \n", expBD09toBD09MC.Round(l), tp)
	}
}
func TestDB09MctoBD09(t *testing.T) {
	tp := DB09MctoBD09(pBDMC).Round(l)
	if !expDB09MctoBD09.Round(l).Equals(tp) {
		t.Errorf("DB09MctoBD09 err, get:%v exp:%v \n", expDB09MctoBD09.Round(l), tp)
	}
}
func TestToMercator(t *testing.T) {
	tp := ToMercator(p).Round(l)
	if !expToMercator.Round(l).Equals(tp) {
		t.Errorf("ToMercator err, get:%v exp:%v \n", expToMercator.Round(l), tp)
	}
}
func TestFromMercator(t *testing.T) {
	tp := FromMercator(pMC).Round(l)
	if !expFromMercator.Round(l).Equals(tp) {
		t.Errorf("FromMercator err, get:%v exp:%v \n", expFromMercator.Round(l), tp)
	}
}

// ---------------------- Benchmark ----------------------
func BenchmarkBD09toGCJ02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BD09toGCJ02(p)
	}
}

func BenchmarkGCJ02toBD09(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCJ02toBD09(p)
	}
}
func BenchmarkWGS84toGCJ02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WGS84toGCJ02(p)
	}
}
func BenchmarkGCJ02toWGS84(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCJ02toWGS84(p)
	}
}
func BenchmarkBD09toWGS84(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BD09toWGS84(p)
	}
}
func BenchmarkWGS84toBD09(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WGS84toBD09(p)
	}
}
func BenchmarkBD09toBD09MC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BD09toBD09MC(p)
	}
}
func BenchmarkDB09MctoBD09(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DB09MctoBD09(pBDMC)
	}
}
func BenchmarkToMercator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToMercator(p)
	}
}
func BenchmarkFromMercator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FromMercator(pMC)
	}
}
