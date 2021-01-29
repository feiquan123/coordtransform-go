package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	ct "github.com/feiquan123/coordtransform"
)

var (
	tip = `坐标转换
coordtransform -t [type] -n [accuracy] [coords split by semicolon]

支持坐标类型:
	WGS84		WGS84坐标系[大地坐标系]
	GCJ02		火星坐标系
	BD09		百度坐标系
	BD09MC		百度墨卡托坐标系
	Mercator 	普通墨卡托坐标系坐标系

Args:
	-h 			help
	-t 			transform type:
		BD09toGCJ02	百度坐标系 -> 火星坐标系
		GCJ02toBD09 	火星坐标系 -> 百度坐标系 	(默认)
		WGS84toGCJ02 	WGS84坐标系 -> 火星坐标系
		GCJ02toWGS84 	火星坐标系 -> WGS84坐标系
		BD09toWGS84 	百度坐标系 -> WGS84坐标系
		WGS84toBD09 	WGS84坐标系 -> 百度坐标系
		BD09toBD09MC 	百度坐标系 -> 百度墨卡托坐标系
		DB09MctoBD09 	百度墨卡托坐标系 -> 百度坐标系
		ToMercator	任意坐标系 -> 墨卡托坐标系
		FromMercator 	墨卡托坐标系 -> 任意坐标系

examples:
	- 百度墨卡托坐标系 -> 百度坐标系
	coordtransform  -t "DB09MctoBD09"  "12732754.092201,3549486.938401;12732775.608553,3549592.609005;"
	
	- 百度墨卡托坐标系 -> 百度坐标系 -> 火星坐标系
	echo "12732754.092201,3549486.938401;12732775.608553,3549592.609005;" | coordtransform  -t "DB09MctoBD09" | coordtransform  -t "BD09toGCJ02" 
`

	t = flag.String("t", "GCJ02toBD09", "coords transform type")
	n = flag.Int("n", 7, "accuracy of float64")
	f = ct.GCJ02toBD09
)

func init() {
	flag.Usage = func() {
		fmt.Println(tip)
		os.Exit(0)
	}

	flag.Parse()

	switch *t {
	case "BD09toGCJ02":
		f = ct.BD09toGCJ02
	case "GCJ02toBD09":
		f = ct.GCJ02toBD09
	case "WGS84toGCJ02":
		f = ct.WGS84toGCJ02
	case "GCJ02toWGS84":
		f = ct.GCJ02toWGS84
	case "BD09toWGS84":
		f = ct.BD09toWGS84
	case "WGS84toBD09":
		f = ct.WGS84toBD09
	case "BD09toBD09MC":
		f = ct.BD09toBD09MC
	case "DB09MctoBD09":
		f = ct.DB09MctoBD09
	case "ToMercator":
		f = ct.ToMercator
	case "FromMercator":
		f = ct.FromMercator
	default:
		f = ct.GCJ02toBD09
	}
}

func transform(ps string, l int) string {
	suffix := false
	if ps[len(ps)-1] == ';' {
		suffix = true
	}

	d := strings.Split(ps, ";")
	data := make([]string, 0, len(d))
	for _, v := range d {
		arr := strings.Split(v, ",")
		if len(arr) != 2 {
			continue
		}
		lon, err := strconv.ParseFloat(arr[0], 64)
		if err != nil {
			panic(err)
		}
		lat, err := strconv.ParseFloat(arr[1], 64)
		if err != nil {
			panic(err)
		}
		data = append(data, f(ct.Point{Lon: lon, Lat: lat}).RoundString(l))
	}
	s := strings.Join(data, ";")
	if suffix {
		s += ";"
	}
	return s
}

func main() {
	if flag.NArg() > 0 {
		fmt.Println(transform(flag.Arg(0), *n))
	} else {
		d, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		fmt.Println(transform(string(d[:len(d)-1]), *n))
	}
}
