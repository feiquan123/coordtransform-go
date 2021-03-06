## coordtransform

coord transform

This module provides coords transform function.

```txt
支持坐标类型:
	WGS84           WGS84坐标系[大地坐标系]
	GCJ02           火星坐标系
	BD09            百度坐标系
	BD09MC          百度墨卡托坐标系
	Mercator        普通墨卡托坐标系坐标系
```

源码:[coordtransform-go](https://github.com/feiquan123/coordtransform-go)

## example

### go

[coordtransform.go](./example/coordtransform.go)

### shell

```sh
# 帮助
coordtransform -h

# 百度墨卡托坐标系 -> 百度坐标系
coordtransform  -t "DB09MctoBD09"  "12732754.092201,3549486.938401;12732775.608553,3549592.609005;"

# 通过管道,百度墨卡托坐标系 -> 百度坐标系 -> 火星坐标系
echo "12732754.092201,3549486.938401;12732775.608553,3549592.609005;" | coordtransform  -t "DB09MctoBD09" | coordtransform  -t "BD09toGCJ02" 
```

## methods

```go
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
```

## install

With [go](https://golang.google.cn/) do:

```sh
go get github.com/feiquan123/coordtransform-go
go install github.com/feiquan123/coordtransform-go/cmd/coordtransform
```

## license

[MIT](https://choosealicense.com/licenses/mit/#)

## refer

百度源码:

- [public.js](http://api.map.baidu.com/lbsapi/getpoint/Js/public.js?20200211)
- [getscript 最新更新时间 20210126](https://api.map.baidu.com/getscript?v=2.0&ak=E4805d16520de693a3fe707cdc962045&services=&t=20210113094335)
