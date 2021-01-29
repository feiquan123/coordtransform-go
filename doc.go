// Package coordtransform 坐标转化
package coordtransform

/*
查看百度源码：
	BaiDu js: https://api.map.baidu.com/getscript?v=2.0&ak=E4805d16520de693a3fe707cdc962045&services=&t=20210113094335
	BaiDu js version: 104
	Baidu js updateDate: 20210126

方法:
|function	|coordtransform|
|-----------|--------------|
|qc			|火星坐标系->百度坐标系|
|BC 		|百度坐标系->火星坐标系|
|T.ub 		|百度墨卡托坐标系 - > 百度坐标系|
|T.tb 		|百度坐标系-> 百度墨卡托坐标系|


支持的全部方法:
	BD09toGCJ02     百度坐标系 -> 火星坐标系
	GCJ02toBD09     火星坐标系 -> 百度坐标系        (默认)
	WGS84toGCJ02    WGS84坐标系 -> 火星坐标系
	GCJ02toWGS84    火星坐标系 -> WGS84坐标系
	BD09toWGS84     百度坐标系 -> WGS84坐标系
	WGS84toBD09     WGS84坐标系 -> 百度坐标系
	BD09toBD09MC    百度坐标系 -> 百度墨卡托坐标系
	DB09MctoBD09    百度墨卡托坐标系 -> 百度坐标系
	ToMercator      任意坐标系 -> 墨卡托坐标系
	FromMercator    墨卡托坐标系 -> 任意坐标系
*/
