// Given civil date/time and geographical longitude calculates Sidereal time.
// Usage:
//
//	lst --datetime=DATETIME --longitude=LONGITUDE
//
// DATETIME is a civil date in RFC3339 format, e.g. 2023-04-13T06:00:00Z
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ilbagatto/go-astro/deltat"
	"github.com/ilbagatto/go-astro/julian"
	"github.com/ilbagatto/go-astro/mathutils"
	"github.com/ilbagatto/go-astro/nutequ"
	"github.com/ilbagatto/go-astro/sidereal"
)

func main() {
	now := time.Now().Format("2024-03-02T21:08:25Z")
	dateStr := flag.String("datetime", now, "date/time in RFC3339 format without time zone")
	lng := flag.Float64("longitude", 0.0, "geographical longitude, negative westwards")
	flag.Parse()

	jd, error := julian.DateStringToJulian(*dateStr)
	if error != nil {
		fmt.Printf("Invalid date: %s\n. Please, use format: y-mm-ddThh:mm:ssZ", *dateStr)
		os.Exit(1)
	}
	dt := deltat.DeltaT(jd)
	jde := jd + dt/86400 // Dynamic time.

	dpsi, deps := nutequ.Nutation(jde)
	eps := nutequ.TrueObliquity(jde, deps)
	opts := sidereal.SiderealOptions{Lng: *lng, Dpsi: dpsi, Eps: eps}
	lst := sidereal.JulianToSidereal(jd, opts)
	hrs, min, sec := mathutils.Hms(lst)

	fmt.Printf("%02d:%02d:%04.1f\n", hrs, min, sec)
}
