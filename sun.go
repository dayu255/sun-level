package sun

import (
	"math"
	"time"
)

// degree to radian
func rad(angle float64) float64 {
	return angle * math.Pi / 180
}

func deg(rad float64) float64 {
	return rad / math.Pi * 180
}

const r2345 float64 = 0.40927971 // 23.45(degrees) * pi / 180
const r360365 float64 = 2 * math.Pi / 365 // 360(degrees) / 365

// 赤緯を求める
func calDelta(t time.Time) float64 {
	var day float64 = float64(t.YearDay())
	return r2345 * math.Sin(r360365 * (day - 81))
}

// 均時差を求める
func calE(t time.Time) float64 {
	B := r360365 * float64(t.YearDay() - 81)
	return 9.87 * math.Sin(2 * B) - 7.53 * math.Cos(B) - 1.5 * math.Sin(B)
}

// 時角を求める
func calHourAngle(t time.Time, lon float64, lat float64, E float64) float64 {
	utc := t.UTC()
	tsUTC := float64(utc.Hour()) + float64(utc.Minute()) / 60.0 + float64(utc.Second()) / 3600.0
	ts := tsUTC + (lon / 15) + (E / 60)
	return rad(15 * (ts - 12))
}

func CalSunLevel(t time.Time, lat float64, lon float64) float64 {
	delta := calDelta(t)
	e := calE(t)
	hourAngle := calHourAngle(t, lon, lat, e)
	h := math.Sin(rad(lat)) * math.Sin(delta) + math.Cos(rad(lat)) * math.Cos(delta) * math.Cos(hourAngle)
	return math.Asin(h)
}