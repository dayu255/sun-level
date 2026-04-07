package main

import (
	"time"
	"fmt"
	"math"
	"github.com/dayu255/sun-level"
)

func main() {
	t := time.Now()
	t = t.Add(-7.0 * time.Hour)
	lat := 35.16
	lon := 136.9
	p := sun.CalSunLevel(t, lat, lon) / math.Pi * 180
	fmt.Println(t)
	fmt.Println(p)
}