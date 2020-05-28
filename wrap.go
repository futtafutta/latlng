package lnglat

// CAD座標用ラップ関数
// ZMAP TOWN2 は日本測地系（旧測地系:Tokyo97）のため変換式をラップする。

// WrapLL2XY ... Zmap(旧測地系)データの経度,緯度=>X,Y変換。ラップ関数
func WrapLL2XY(lng, lat float64, kei uint8) (x, y float64) {
	// WSG84 => 日本測地系 変換
	lng, lat = ConvGeodeticDatum(lng, lat, ModeConvLL2XY, ModeGeoWorld2Tokyo)
	x, y = LL2XY(lng, lat, kei)
	return
}

// WrapXY2LL ... Zmap(旧測地系)データのX,Y=>経度,緯度変換。ラップ関数
func WrapXY2LL(x, y float64, kei uint8) (lng, lat float64) {
	lng, lat = XY2LL(x, y, kei)
	// 日本測地系 => WSG84 変換
	lng, lat = ConvGeodeticDatum(lng, lat, ModeConvXY2LL, ModeGeoTokyo2World)
	//x, y = LL2XY(lng, lat, kei)
	return
}

// 変換モードの指定。LL2XY or XY2LL
const (
	// ModeConvLL2XY ... 経度,緯度 => X,Y
	ModeConvLL2XY = iota
	// ModeConvXY2LL ... X,Y => 経度,緯度
	ModeConvXY2LL
)

// 変換測地系の指定。日本測地系 or WGS84
const (
	// ModeGeoTokyo2World ... 日本測地系 => WGS84
	ModeGeoTokyo2World = iota
	// ModeGeoWorld2Tokyo ... WGS84 => 日本測地系
	ModeGeoWorld2Tokyo
)

// ConvGeodeticDatum ... 測地系相互変換関数。日本測地系 <--> WGS84
//  引数 = oB:経度, oL:緯度
// 	戻り値 = lng:経度, lat:緯度
func ConvGeodeticDatum(oB, oL float64, modeconv, modegeo uint) (lng, lat float64) {
	switch modegeo {
	case ModeGeoTokyo2World:
		lat = oB - 0.00010695*oB + 0.000017464*oL + 0.0046017
		lng = oL - 0.000046038*oB - 0.000083043*oL + 0.010040
	case ModeGeoWorld2Tokyo:
		oB, oL = oL, oB
		lat = oB + 0.00010696*oB - 0.000017467*oL - 0.0046020
		lng = oL + 0.000046047*oB + 0.000083049*oL - 0.010041
	}

	return
}
