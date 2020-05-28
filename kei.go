package lnglat

import "fmt"

const (
	// KeiMin ... 19座標系最小値
	KeiMin = 1
	// KeiMax ... 19座標系最大値
	KeiMax = 19
)

// Origin ... 19座標系原点情報
type Origin struct {
	Kei uint8
	*LngLat
	Srs string
}

func newOrigin(kei uint8, lng, lat float64, srs string) (o *Origin) {
	o = new(Origin)
	o.Kei = kei
	ll := new(LngLat)
	ll.Longitude = lng
	ll.Latitude = lat
	o.LngLat = ll
	o.Srs = srs
	return o
}

// 緯度経度型MAP
type originMap map[uint8]*Origin

// GetKeiOrigin ... 原点座標位置を戻す
func GetKeiOrigin(kei uint8) (o *Origin, err error) {
	omap := makeKeiOrigin()
	o, ok := omap[kei]
	if !ok {
		err = fmt.Errorf("*** Error *** 座標系不正（%d^%dで指定して下さい。）", KeiMin, KeiMax)
	}
	return
}

func makeKeiOrigin() (omap originMap) {
	omap = make(originMap)
	omap[1] = newOrigin(1, 129.5, 33.0, "+proj=tmerc +lat_0=33 +lon_0=129.5 +k=0.9999 +x_0=0 +y_0=0 +ellps=bessel +towgs84=-146.414,507.337,680.507,0,0,0,0 +units=m +no_defs")
	omap[2] = newOrigin(2, 131.0, 33.0, "+proj=tmerc +lat_0=33 +lon_0=131 +k=0.9999 +x_0=0 +y_0=0 +ellps=bessel +towgs84=-146.414,507.337,680.507,0,0,0,0 +units=m +no_defs")
	omap[3] = newOrigin(3, 132.166666666667, 36.0, "+proj=tmerc +lat_0=36 +lon_0=132.166666666667 +k=0.9999 +x_0=0 +y_0=0 +ellps=bessel +towgs84=-146.414,507.337,680.507,0,0,0,0 +units=m +no_defs")
	omap[4] = newOrigin(4, 133.5, 33.0, "+proj=tmerc +lat_0=33 +lon_0=133.5 +k=0.9999 +x_0=0 +y_0=0 +ellps=bessel +towgs84=-146.414,507.337,680.507,0,0,0,0 +units=m +no_defs")
	omap[5] = newOrigin(5, 134.333333333333, 36.0, "+proj=tmerc +lat_0=36 +lon_0=134.333333333333 +k=0.9999 +x_0=0 +y_0=0 +ellps=bessel +towgs84=-146.414,507.337,680.507,0,0,0,0 +units=m +no_defs")
	omap[6] = newOrigin(6, 136.0, 36.0, "+proj=tmerc +lat_0=36 +lon_0=136 +k=0.9999 +x_0=0 +y_0=0 +ellps=bessel +towgs84=-146.414,507.337,680.507,0,0,0,0 +units=m +no_defs")
	omap[7] = newOrigin(7, 137.166666666667, 36.0, "+proj=tmerc +lat_0=36 +lon_0=137.166666666667 +k=0.9999 +x_0=0 +y_0=0 +ellps=bessel +towgs84=-146.414,507.337,680.507,0,0,0,0 +units=m +no_defs")
	omap[8] = newOrigin(8, 138.5, 36.0, "+proj=tmerc +lat_0=36 +lon_0=138.5 +k=0.9999 +x_0=0 +y_0=0 +ellps=bessel +towgs84=-146.414,507.337,680.507,0,0,0,0 +units=m +no_defs")
	omap[9] = newOrigin(9, 139.833333333333, 36.0, "+proj=tmerc +lat_0=36 +lon_0=139.833333333333 +k=0.9999 +x_0=0 +y_0=0 +ellps=bessel +towgs84=-146.414,507.337,680.507,0,0,0,0 +units=m +no_defs")
	omap[10] = newOrigin(10, 140.833333333333, 40.0, "+proj=tmerc +lat_0=40 +lon_0=140.833333333333 +k=0.9999 +x_0=0 +y_0=0 +ellps=bessel +towgs84=-146.414,507.337,680.507,0,0,0,0 +units=m +no_defs")
	omap[11] = newOrigin(11, 140.25, 44.0, "+proj=tmerc +lat_0=44 +lon_0=140.25 +k=0.9999 +x_0=0 +y_0=0 +ellps=bessel +towgs84=-146.414,507.337,680.507,0,0,0,0 +units=m +no_defs")
	omap[12] = newOrigin(12, 142.25, 44.0, "+proj=tmerc +lat_0=44 +lon_0=142.25 +k=0.9999 +x_0=0 +y_0=0 +ellps=bessel +towgs84=-146.414,507.337,680.507,0,0,0,0 +units=m +no_defs")
	omap[13] = newOrigin(13, 144.25, 44.0, "+proj=tmerc +lat_0=44 +lon_0=144.25 +k=0.9999 +x_0=0 +y_0=0 +ellps=bessel +towgs84=-146.414,507.337,680.507,0,0,0,0 +units=m +no_defs")
	omap[14] = newOrigin(14, 142, 26.0, "+proj=tmerc +lat_0=26 +lon_0=142 +k=0.9999 +x_0=0 +y_0=0 +ellps=bessel +towgs84=-146.414,507.337,680.507,0,0,0,0 +units=m +no_defs")
	omap[15] = newOrigin(15, 127.5, 26.0, "+proj=tmerc +lat_0=26 +lon_0=127.5 +k=0.9999 +x_0=0 +y_0=0 +ellps=bessel +towgs84=-146.414,507.337,680.507,0,0,0,0 +units=m +no_defs")
	omap[16] = newOrigin(16, 124.0, 26.0, "+proj=tmerc +lat_0=26 +lon_0=124 +k=0.9999 +x_0=0 +y_0=0 +ellps=bessel +towgs84=-146.414,507.337,680.507,0,0,0,0 +units=m +no_defs")
	omap[17] = newOrigin(17, 131.0, 26.0, "+proj=tmerc +lat_0=26 +lon_0=131 +k=0.9999 +x_0=0 +y_0=0 +ellps=bessel +towgs84=-146.414,507.337,680.507,0,0,0,0 +units=m +no_defs")
	omap[18] = newOrigin(18, 136.0, 20.0, "+proj=tmerc +lat_0=20 +lon_0=136 +k=0.9999 +x_0=0 +y_0=0 +ellps=bessel +towgs84=-146.414,507.337,680.507,0,0,0,0 +units=m +no_defs")
	omap[19] = newOrigin(19, 154.0, 26.0, "+proj=tmerc +lat_0=26 +lon_0=154 +k=0.9999 +x_0=0 +y_0=0 +ellps=bessel +towgs84=-146.414,507.337,680.507,0,0,0,0 +units=m +no_defs")

	return
}
