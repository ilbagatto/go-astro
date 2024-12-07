package nutequ

import (
	"testing"

	"github.com/ilbagatto/go-astro/julian"
	"github.com/ilbagatto/go-astro/mathutils"
)

type _NutTestCase struct {
	jd   float64
	dpsi float64
	deps float64
}

var nutCases = [...]_NutTestCase{
	{
		jd:   2399215.5, // 1856 Sept. 23
		dpsi: -0.00127601021242336,
		deps: 0.00256293723137559,
	},
	{
		jd:   2451544.5, // 2000 Jan. 1
		dpsi: -0.00387728730373955,
		deps: -0.00159919822661103,
	},
	{
		jd:   2443825.69, // 1978 Nov 17
		dpsi: -9.195562346652888e-04,
		deps: -2.635113483663831e-03,
	},
	{
		jd:   julian.CivilToJulian(julian.CivilDate{Year: 1989, Month: 2, Day: 4}),
		dpsi: 0.0023055555555555555,
		deps: 0.0022944444444444444,
	},
	{
		jd:   julian.CivilToJulian(julian.CivilDate{Year: 2000, Month: 1, Day: 1.5}),
		dpsi: -0.003877777777777778,
		deps: -0.0016,
	},
	{
		jd:   julian.CivilToJulian(julian.CivilDate{Year: 1995, Month: 4, Day: 23}),
		dpsi: 0.0026472222222222223,
		deps: -0.002013888888888889,
	},
	{
		jd:   julian.CivilToJulian(julian.CivilDate{Year: 1965, Month: 2, Day: 1}),
		dpsi: -0.0042774118548615766,
		deps: 0.000425,
	},
}

func TestNutLon(t *testing.T) {
	for _, test := range nutCases {
		dpsi, _ := Nutation(test.jd)
		if !mathutils.AlmostEqual(dpsi, test.dpsi, 1e-4) {
			t.Errorf("Expected: %f, got: %f", test.dpsi, dpsi)
		}
	}
}

func TestNutEcl(t *testing.T) {
	for _, test := range nutCases {
		_, deps := Nutation(test.jd)
		if !mathutils.AlmostEqual(deps, test.deps, 1e-4) {
			t.Errorf("Expected: %f, got: %f", test.deps, deps)
		}
	}
}
