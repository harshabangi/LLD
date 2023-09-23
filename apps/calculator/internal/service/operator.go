package service

type ArithmeticOperator interface {
	Evaluate(op1, op2 float64) float64
	Precedence() int
}

type Addition struct{}

func (a Addition) Evaluate(op1, op2 float64) float64 {
	return op1 + op2
}

func (a Addition) Precedence() int {
	return 1
}

type Subtraction struct{}

func (s Subtraction) Evaluate(op1, op2 float64) float64 {
	return op1 - op2
}

func (s Subtraction) Precedence() int {
	return 1
}

type Multiplication struct{}

func (m Multiplication) Evaluate(op1, op2 float64) float64 {
	return op1 * op2
}

func (m Multiplication) Precedence() int {
	return 2
}

type Division struct{}

func (d Division) Evaluate(op1, op2 float64) float64 {
	if op2 == 0 {
		panic("denominator should not be zero")
	}
	return op1 / op2
}

func (d Division) Precedence() int {
	return 2
}

func arithmeticOperatorFactory(operator string) ArithmeticOperator {
	switch operator {
	case "+":
		return Addition{}
	case "-":
		return Subtraction{}
	case "*":
		return Multiplication{}
	case "/":
		return Division{}
	default:
		return nil
	}
}
