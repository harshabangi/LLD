package service

import (
	"github.com/harshabangi/LLD/apps/calculator/internal/memory"
	"github.com/harshabangi/LLD/apps/calculator/pkg"
	"strconv"
	"time"
)

type CalculatorService struct {
	memoryStore memory.Store
}

func NewCalculatorService() *CalculatorService {
	return &CalculatorService{
		memoryStore: memory.NewMemoryStore(),
	}
}

func (s *CalculatorService) Calculate(expression string, recordName string) float64 {

	result := evaluate(expression)
	now := time.Now()

	calculationRecord := memory.CalculationRecord{
		Name:        recordName,
		Expression:  expression,
		Result:      result,
		SubmittedAt: &now,
	}
	s.memoryStore.AddCalculation(calculationRecord)

	return result
}

func evaluate(infixExpression string) float64 {
	tokens := infixToPostfix(infixExpression)
	return reversePolishNotation(tokens)
}

func (s *CalculatorService) ListCalculations() []memory.CalculationRecord {
	return s.memoryStore.ListCalculationRecords()
}

func (s *CalculatorService) GetCalculationRecord(recordName string) pkg.CalculationRecord {
	if recordName != "" {
		record := s.memoryStore.GetCalculationRecord(recordName)
		return record.ToModel()
	}
	return pkg.CalculationRecord{}
}

func reversePolishNotation(tokens []string) float64 {
	sta := Stack[float64]{}

	for _, token := range tokens {
		operator := arithmeticOperatorFactory(token)

		if operator == nil {
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				panic(err)
			}
			sta.Push(num)
		} else {
			operand1, operand2 := *sta.Pop(), *sta.Pop()
			sta.Push(operator.Evaluate(operand1, operand2))
		}
	}
	return *sta.Top()
}

func infixToPostfix(infix string) []string {
	var sta Stack[string]
	var postfix []string

	for _, char := range infix {
		opChar := string(char)

		switch {
		case char >= '0' && char <= '9':
			postfix = append(postfix, opChar)
		case char == '(':
			sta.Push(opChar)
		case char == ')':
			for *sta.Top() != "(" {
				postfix = append(postfix, *sta.Top())
				sta.Pop()
			}
			sta.Pop()
		default:
			for !sta.IsEmpty() &&
				arithmeticOperatorFactory(opChar).Precedence() <= arithmeticOperatorFactory(*sta.Top()).Precedence() {
				postfix = append(postfix, *sta.Top())
				sta.Pop()
			}
			sta.Push(opChar)
		}
	}

	for !sta.IsEmpty() {
		postfix = append(postfix, *sta.Top())
		sta.Pop()
	}

	return postfix
}
