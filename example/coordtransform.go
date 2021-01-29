package main

import (
	"fmt"

	ct "github.com/feiquan123/coordtransform"
)

func main() {
	p := ct.Point{114.363089, 30.517411}
	pBDMC := ct.Point{12730979.361410993, 3548336.555792139}
	pMC := ct.Point{12730840.83125368, 3569413.5156310988}

	// BD09toGCJ02: 百度坐标系->火星坐标系
	fmt.Println("BD09toGCJ02:", ct.BD09toGCJ02(p))
	// GCJ02toBD09: 火星坐标系->百度坐标系
	fmt.Println("GCJ02toBD09:", ct.GCJ02toBD09(p))
	// WGS84toGCJ02: 火星坐标系->WGS84坐标系
	fmt.Println("WGS84toGCJ02:", ct.WGS84toGCJ02(p))
	// GCJ02toWGS84 火星坐标系->WGS84坐标系
	fmt.Println("GCJ02toWGS84:", ct.GCJ02toWGS84(p))
	// BD09toWGS84 百度坐标系->WGS84坐标系
	fmt.Println("BD09toWGS84:", ct.BD09toWGS84(p))
	// WGS84toBD09 WGS84坐标系->百度坐标系
	fmt.Println("WGS84toBD09:", ct.WGS84toBD09(p))
	// BD09toBD09MC 百度坐标系-> 百度墨卡托坐标系
	fmt.Println("BD09toBD09MC:", ct.BD09toBD09MC(p))
	// DB09MctoBD09 百度墨卡托坐标系 - > 百度坐标
	fmt.Println("DB09MctoBD09:", ct.DB09MctoBD09(pBDMC))
	// ToMercator 任意坐标系-> 墨卡托坐标
	fmt.Println("ToMercator:", ct.ToMercator(p))
	// FromMercator 墨卡托坐标系-> 任意坐标系
	fmt.Println("FromMercator:", ct.FromMercator(pMC))
}
