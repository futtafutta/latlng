//============================================================================
// 経度緯度⇔平面座標 変換パッケージ
//============================================================================

package lnglat

import (
	"fmt"
	"math"
)

// XY緯度経度変換式定数
const (
	ra = 6377397.155
	rb = 6356078.965
	k1 = 1.00503730604855
	k2 = 5.0478492403e-3
	k3 = 1.0563786831e-5
	k4 = 2.0633322e-8
	m0 = 0.9999
)

// LL2XY 緯度経度を平面座標（測地座標）に変換
// 引数 = lng:経度, lat:緯度, kei :19座標系の系番号
// 戻り値 =  x:19座標系のX座標, var y:19座標系のY座標
//
func LL2XY(lng, lat float64, kei uint8) (x, y float64) {

	org, err := GetKeiOrigin(kei)
	if err != nil {
		fmt.Println(err)
		return
	}
	b0p, L0p := org.Latitude, org.Longitude

	b := lat
	L := lng
	b0 := b0p
	L0 := L0p

	r0 := 180 / math.Pi
	b0 = b0 / r0
	L0 = L0 / r0
	b = b / r0
	L = L / r0

	e := math.Sqrt(1 - math.Pow((rb/ra), 2))
	e1 := math.Sqrt(math.Pow((ra/rb), 2) - 1)
	w := math.Sqrt(1 - math.Pow((e*math.Sin(b)), 2))
	n := ra / w
	a1 := k1 * (b - b0)
	a2 := k2 * (math.Sin(2*b) - math.Sin(2*b0)) / 2
	a3 := k3 * (math.Sin(4*b) - math.Sin(4*b0)) / 4
	a4 := k4 * (math.Sin(6*b) - math.Sin(6*b0)) / 6
	db := ra * (1 - math.Pow(e, 2)) * (a1 - a2 + a3 - a4)
	t := math.Tan(b)
	h := e1 * math.Cos(b)
	dl := L - L0

	y = db + math.Pow(dl, 2)*n/2*math.Sin(b)*math.Cos(b)
	y = y + math.Pow(dl, 4)*n/24*math.Sin(b)*math.Pow(math.Cos(b), 3)*(5-math.Pow(t, 2)+9*math.Pow(h, 2)+4*math.Pow(h, 4))
	y = y + math.Pow(dl, 6)*n/720*math.Sin(b)*math.Pow(math.Cos(b), 5)*(61-58*math.Pow(t, 2)+math.Pow(t, 4)+270*math.Pow(h, 2)-330*math.Pow((h*t), 2))
	x = dl*n*math.Cos(b) + math.Pow(dl*math.Cos(b), 3)*n/6*(1-math.Pow(t, 2)+math.Pow(h, 2))
	x = x + math.Pow((dl*math.Cos(b)), 5)*n/120*(5-18*math.Pow(t, 2)+math.Pow(t, 4)+14*math.Pow(h, 2)-58*math.Pow((h*t), 2))

	x *= m0
	y *= m0

	return

}

// XY2LL 平面座標（測地座標）を緯度経度に変換する
// 引数 = x:19座標系のX座標, y:19座標系のY座標, kei :19座標系の系番号
// 戻り値 = lng:経度, lat:緯度
//
func XY2LL(x, y float64, kei uint8) (lng, lat float64) {

	org, err := GetKeiOrigin(kei)
	if err != nil {
		fmt.Println(err)
		return
	}
	b0p, L0p := org.Latitude, org.Longitude

	b0 := b0p
	L0 := L0p

	r0 := 180 / math.Pi
	b0 = b0 / r0
	L0 = L0 / r0

	e := math.Sqrt(1 - math.Pow((rb/ra), 2))
	e1 := math.Sqrt(math.Pow((ra/rb), 2) - 1)
	a := ra * (1 - math.Pow(e, 2))
	d := a*(k1*b0-k2/2*math.Sin(2*b0)+k3/4*math.Sin(4*b0)-k4/6*math.Sin(6*b0)) + y/m0
	p := make([]float64, 7)
	s := make([]float64, 7)
	m := make([]float64, 7)
	for i := 1; i <= 6; i++ {
		if i == 1 {
			p[i] = b0
		} else {
			p[i] = p[i-1] - (s[i-1]-d)/m[i-1]
		}
		s[i] = a * (k1*p[i] - k2/2*math.Sin(2*p[i]) + k3/4*math.Sin(4*p[i]) - k4/6*math.Sin(6*p[i]))
		m[i] = a / (math.Pow((1 - math.Pow((e*math.Sin(p[i])), 2)), 1.5))
	}
	pp := p[5]
	x1 := x / m0
	t := math.Tan(pp)
	h := e1 * math.Cos(pp)
	n := ra / (math.Pow((1 - math.Pow((e*math.Sin(pp)), 2)), 0.5))
	mm := m[5]

	LB := pp - (math.Pow(x1, 2)*t/(2*mm*n) - math.Pow(x1, 4)*t*(5+3*math.Pow(t, 2)+math.Pow(h, 2)-9*math.Pow((t*h), 2))/(24*mm*math.Pow(n, 3)))
	LB = LB - (math.Pow(x1, 6) * t * (61 + 90*math.Pow(t, 2) + 45*math.Pow(t, 4)) / (720 * mm * math.Pow(n, 5)))

	L := (x1/(n*math.Cos(pp)) - math.Pow(x1, 3)*(1+2*math.Pow(t, 2)+math.Pow(h, 2))/(6*math.Pow(n, 3)*math.Cos(pp)))
	L = L + (math.Pow(x1, 5) * (5 + 28*math.Pow(t, 2) + 24*math.Pow(t, 4)) / (120 * math.Pow(n, 5) * math.Cos(pp)))
	L = L + L0

	L = L * r0
	LB = LB * r0
	lng = LB
	lat = L

	return
}
