package mathutils

import "testing"

func TestAlmostEqualPositive(t *testing.T) {
	if !AlmostEqual(10.1234567, 10.1234568, 1e-6) {
		t.Errorf("Numbers should be considered equal")
	}
}

func TestAlmostEqualNegative(t *testing.T) {
	if AlmostEqual(10.123457, 10.123456, 1e-6) {
		t.Errorf("Numbers should not be considered equal")
	}
}

func TestPositiveFrac(t *testing.T) {
	if !AlmostEqual(Frac(5.5), 0.5, 1e-6) {
		t.Errorf("frac(5.5) should be 0.5")
	}
}

func TestNegativeFrac(t *testing.T) {
	if !AlmostEqual(Frac(-5.5), -0.5, 1e-6) {
		t.Errorf("frac(-5.5) should be -0.5")
	}
}

func TestReduceHoursPositive(t *testing.T) {
	if !AlmostEqual(ReduceHours(49.5), 1.5, 1e-6) {
		t.Errorf("49.5 should be reduced to 1.5")
	}
}

func TestReduceHoursNegative(t *testing.T) {
	if !AlmostEqual(ReduceHours(-0.5), 23.5, 1e-6) {
		t.Errorf("-0.5 should be reduced to 23.5")
	}
}

func TestReduceDegPositive(t *testing.T) {
	if !AlmostEqual(ReduceDeg(324070.45), 70.45, 1e-6) {
		t.Errorf("324070.45 should be reduced to 70.45")
	}
}

func TestReduceDegNegative(t *testing.T) {
	if !AlmostEqual(ReduceHours(-700), 20, 1e-6) {
		t.Errorf("-700 should be reduced to 20")
	}
}

func TestReduceRadPositive(t *testing.T) {
	if !AlmostEqual(ReduceRad(12.89), 0.323629385640829, 1e-6) {
		t.Errorf("12.89 should be reduced to 0.323629385640829")
	}
}

func TestReduceRadNegative(t *testing.T) {
	if !AlmostEqual(ReduceRad(-12.89), 5.95955592153876, 1e-6) {
		t.Errorf("-12.89 should be reduced to 5.95955592153876")
	}
}

func TestShortPolynome(t *testing.T) {
	got := Polynome(10, 1, 2, 3)
	exp := 321.0
	if !AlmostEqual(got, exp, 1e-6) {
		t.Errorf("Expected: %f, got: %f", exp, got)
	}
}

func TestLongPolynome(t *testing.T) {
	got := Polynome(
		-0.127296372347707,
		0.409092804222329,
		-0.0226937890431606,
		-7.51461205719781e-06,
		0.0096926375195824,
		-0.00024909726935408,
		-0.00121043431762618,
		-0.000189319742473274,
		3.4518734094999e-05,
		0.000135117572925228,
		2.80707121362421e-05,
		1.18779351871836e-05)
	exp := 0.411961500152426
	if !AlmostEqual(got, exp, 1e-6) {
		t.Errorf("Expected: %f, got: %f", exp, got)
	}
}

func TestFrac360(t *testing.T) {
	const k = 23772.99 / 36525

	var cases = [5][2]float64{
		{31.7842235930254, 1.000021358e2},
		{30.6653235575305, 9.999736056e1},
		{42.3428797768338, 1.336855231e3},
		{273.934866366267, 1.325552359e3},
		{178.873057561472, 5.37261666700},
	}

	for _, test := range cases {
		exp := test[0]
		arg := test[1] * k
		got := Frac360(arg)
		if !AlmostEqual(got, exp, 1e-6) {
			t.Errorf("Expected: %f, got: %f", exp, got)
		}
	}
}

func TestPositiveSexagesimal(t *testing.T) {
	h, m, s := Hms(20.75833333333333)
	if h != 20 {
		t.Errorf("Expected: %d, got: %d", 20, h)
	}
	if m != 45 {
		t.Errorf("Expected: %d, got: %d", 45, m)
	}
	if !AlmostEqual(s, 30, 1e-6) {
		t.Errorf("Expected: %f, got: %f", 30.0, s)
	}

}

func TestNegativeSexagesimal(t *testing.T) {
	h, m, s := Hms(-20.75833333333333)
	if h != -20 {
		t.Errorf("Expected: %d, got: %d", -20, h)
	}
	if m != 45 {
		t.Errorf("Expected: %d, got: %d", 45, m)
	}
	if !AlmostEqual(s, 30, 1e-6) {
		t.Errorf("Expected: %f, got: %f", 30.0, s)
	}

}
