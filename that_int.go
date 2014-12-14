package assert

func ThatInt(t TestingT, actual int) *IntAssert {
	return &IntAssert{t, actual}
}

type IntAssert struct {
	t      TestingT
	actual int
}

func (assert *IntAssert) isTrue(condition bool, format string, args ...interface{}) *IntAssert {
	if !condition {
		assert.t.Errorf(format, args...)
	}
	return assert
}

func (assert *IntAssert) IsZero() *IntAssert {
	return assert.isTrue(assert.actual == 0,
		"Expected zero, but was <%d>.", assert.actual)
}

func (assert *IntAssert) IsNonzero() *IntAssert {
	return assert.isTrue(assert.actual != 0,
		"Expected nonzero, but was zero.")
}

func (assert *IntAssert) IsEqualTo(expected int) *IntAssert {
	return assert.isTrue(assert.actual == expected,
		"Expected integer equal to <%d>, but was <%d>.", expected, assert.actual)
}

func (assert *IntAssert) IsNotEqualTo(nonexpected int) *IntAssert {
	return assert.isTrue(assert.actual != nonexpected,
		"Expected integer not equal to <%d>, but was equal.", nonexpected)
}

func (assert *IntAssert) IsPositive() *IntAssert {
	return assert.isTrue(assert.actual > 0,
		"Expected positive integer, but was <%d>.", assert.actual)
}

func (assert *IntAssert) IsNonpositive() *IntAssert {
	return assert.isTrue(assert.actual <= 0,
		"Expected nonpositive integer, but was <%d>.", assert.actual)
}

func (assert *IntAssert) IsNegative() *IntAssert {
	return assert.isTrue(assert.actual < 0,
		"Expected negative integer, but was <%d>.", assert.actual)
}

func (assert *IntAssert) IsNonnegative() *IntAssert {
	return assert.isTrue(assert.actual >= 0,
		"Expected nonnegative integer, but was <%d>.", assert.actual)
}

func (assert *IntAssert) IsGreaterThan(lesserOrEqual int) *IntAssert {
	return assert.isTrue(assert.actual > lesserOrEqual,
		"Expected integer greater than <%d>, but was <%d>.", lesserOrEqual, assert.actual)
}

func (assert *IntAssert) IsGreaterOrEqualTo(lesser int) *IntAssert {
	return assert.isTrue(assert.actual >= lesser,
		"Expected integer greater or equal to <%d>, but was <%d>.", lesser, assert.actual)
}

func (assert *IntAssert) IsLessThan(greaterOrEqual int) *IntAssert {
	return assert.isTrue(assert.actual < greaterOrEqual,
		"Expected integer lesser than <%d>, but was <%d>.", greaterOrEqual, assert.actual)
}

func (assert *IntAssert) IsLessOrEqualTo(greater int) *IntAssert {
	return assert.isTrue(assert.actual <= greater,
		"Expected integer lesser or equal to <%d>, but was <%d>.", greater, assert.actual)
}
