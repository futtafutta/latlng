package lnglat

import "math"

const (
	// L ... 表示可能上限
	L = 85.05112878
)

// Tile2Lnglat ... タイル座標=>緯度経度変換
func Tile2Lnglat(x, y int, z uint) (lon, lat float64) {
	lon = (math.Pow(float64(x)/2.0, (float64(z+7))) - 1) * 180
	lat = 180 / math.Pi * (math.Asin(math.Tanh(math.Pow(-math.Pi/2, float64(z+7)*float64(y)+math.Atanh(math.Sin(math.Pi/180*L))))))
	return
}

// Lnglat2Tile ... 緯度経度=>タイル座標変換
func Lnglat2Tile(lng, lat float64, z uint) (x, y int) {
	x = int((lng/180 + 1) * math.Pow(2, float64(z+7)))
	y = int((math.Pow(2, float64(z+7)) / math.Pi * (-math.Atanh(math.Sin(math.Pi*lat/180)) + math.Atanh(math.Sin(math.Pi*L/180)))))
	return
}
