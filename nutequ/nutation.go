// Calculates obliquity of the ecliptic and the effects of nutation on the ecliptic longitude,
// dpsi, and on the obliquity of the ecliptic, deps with accuracy of about 1 arcsecond.
//
// Source: P.Duffett-Smith, "Astronomy with Your PC", 2 edition.
package nutequ

import (
	"math"

	"github.com/ilbagatto/go-astro/julian"
	"github.com/ilbagatto/go-astro/mathutils"
)

// Given a Julian Date, calculate nutation in longitude and obliquity, both in arc-degrees.
func Nutation(jd float64) (dpsi float64, deps float64) {
	t := (jd - julian.J1900) / julian.DAYS_PER_CENT
	t2 := t * t

	ls := 2.796967e2 + 3.030e-4*t2 + mathutils.Frac360(1.000021358e2*t) // Sun's mean longitude
	ms := 3.584758e2 - 1.500e-4*t2 + mathutils.Frac360(9.999736056e1*t) // Sun's mean anomaly
	ms = mathutils.Radians(ms)
	ld := 2.704342e2 - 1.133e-3*t2 + mathutils.Frac360(1.336855231e3*t) // Moon's mean longitude
	md := 2.961046e2 + 9.192e-3*t2 + mathutils.Frac360(1.325552359e3*t) // Moon's mean anomaly
	md = mathutils.Radians(md)
	nm := 2.591833e2 + 2.078e-3*t2 - mathutils.Frac360(5.372616667*t) // longitude of the Moon's ascending node
	nm = mathutils.Radians(nm)
	tls := mathutils.Radians(ls + ls)
	tld := mathutils.Radians(ld + ld)
	tnm := nm + nm

	dpsi = (-17.2327-1.737e-2*t)*math.Sin(nm) +
		(-1.2729-1.3e-4*t)*math.Sin(tls) +
		2.088e-1*math.Sin(tnm) -
		2.037e-1*math.Sin(tld) +
		(1.261e-1-3.1e-4*t)*math.Sin(ms) +
		6.75e-2*math.Sin(md) -
		(4.97e-2-1.2e-4*t)*math.Sin(tls+ms) -
		3.42e-2*math.Sin(tld-nm) -
		2.61e-2*math.Sin(tld+md) +
		2.14e-2*math.Sin(tls-ms) -
		1.49e-2*math.Sin(tls-tld+md) +
		1.24e-2*math.Sin(tls-nm) +
		1.14e-2*math.Sin(tld-md)

	deps = (9.21+9.1e-4*t)*math.Cos(nm) +
		(5.522e-1-2.9e-4*t)*math.Cos(tls) -
		9.04e-2*math.Cos(tnm) +
		8.84e-2*math.Cos(tld) +
		2.16e-2*math.Cos(tls+ms) +
		1.83e-2*math.Cos(tld-nm) +
		1.13e-2*math.Cos(tld+md) -
		9.3e-3*math.Cos(tls-ms) -
		6.6e-3*math.Cos(tls-nm)

	return dpsi / 3600, deps / 3600
}
