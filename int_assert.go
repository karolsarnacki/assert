package assert

type IntAssert interface {
	IsZero() IntAssert
	IsNonZero() IntAssert
	IsEqualTo(expected int) IntAssert
	IsNotEqualTo(unexpected int) IntAssert
	IsPositive() IntAssert
	IsNonPositive() IntAssert
	IsNegative() IntAssert
	IsNonNegative() IntAssert
	IsGreaterThan(less int) IntAssert
	IsGreaterOrEqualTo(lessOrEqual int) IntAssert
	IsLessThan(greater int) IntAssert
	IsLessOrEqualTo(greaterOrEqual int) IntAssert
	IsBetween(lowerBound, upperBound int) IntAssert
	IsNotBetween(lowerBound, upperBound int) IntAssert
	IsIn(elements ...int) IntAssert
	IsNotIn(elements ...int) IntAssert
	IsEven() IntAssert
	IsOdd() IntAssert
	IsDivisibleBy(divisor int) IntAssert
	IsNotDivisibleBy(divisor int) IntAssert
	IsPrime() IntAssert
	IsNotPrime() IntAssert
	IsAnswerToTheUltimateQuestionOfLife() IntAssert
}
