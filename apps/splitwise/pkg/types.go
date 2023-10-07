package pkg

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Split struct {
	UserEmail string `json:"user_email"`
	Amount    int64  `json:"amount"`
}

type PercentSplit struct {
	UserEmail string `json:"user_email"`
	Percent   int64  `json:"percent"`
}

type Expense struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Amount      int64  `json:"amount"`
	PaidBy      string `json:"paid_by"`
}

func (e *Expense) Validate() error {
	if e.Description == "" || e.PaidBy == "" || e.Amount == 0 {
		return fmt.Errorf("inadequate input parameters. Required description, paid_by, amount")
	}
	return nil
}

func (e *Expense) GetPaidBy() string {
	return e.PaidBy
}

type Expenser interface {
	Validate() error
	GetPaidBy() string
	GetSplits() []Split
}

type EqualExpense struct {
	Expense
	Users []string
}

func (e *EqualExpense) Validate() error {
	return e.Expense.Validate()
}

func (e *EqualExpense) GetSplits() []Split {
	n := int64(len(e.Users))
	equalAmount := e.Amount / n

	s := make([]Split, n)

	for i, v := range e.Users {
		s[i] = Split{
			UserEmail: v,
			Amount:    equalAmount,
		}
	}
	return s
}

type ExactExpense struct {
	Expense
	ExactSplits []Split
}

func (e *ExactExpense) GetSplits() []Split {
	return e.ExactSplits
}

func (e *ExactExpense) Validate() error {
	if err := e.Expense.Validate(); err != nil {
		return err
	}

	var expectedAmount int64
	for i := 0; i < len(e.ExactSplits); i++ {
		expectedAmount += e.ExactSplits[i].Amount
	}

	if expectedAmount != e.Amount {
		return fmt.Errorf("something wrong")
	}
	return nil
}

type PercentExpense struct {
	Expense
	percentSplits []PercentSplit
}

func (p *PercentExpense) Validate() error {
	if err := p.Expense.Validate(); err != nil {
		return err
	}

	var total int64 = 0
	for _, v := range p.percentSplits {
		total += v.Percent
	}

	if total == 100 {
		return fmt.Errorf("percent splits doesn't add up to 100")
	}
	return nil
}

func (p *PercentExpense) GetSplits() []Split {
	totalAmount := p.Amount
	splits := make([]Split, len(p.percentSplits))

	for i, v := range p.percentSplits {
		a := (v.Percent / 100) * totalAmount

		splits[i] = Split{
			UserEmail: v.UserEmail,
			Amount:    a,
		}
	}
	return splits
}

func (e *Expense) UnmarshalJSON(data []byte) error {
	m := map[string]interface{}{}

	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	v, ok := m["type"]
	if !ok {
		return fmt.Errorf("no expense type specified in %s", string(data))
	}

	vs, ok := v.(string)
	if !ok {
		return fmt.Errorf("type is not a stirng in %s", string(data))
	}

	switch vs {
	case "equal":
		eqEx := &EqualExpense{}
		if err = json.Unmarshal(data, eqEx); err != nil {
			return err
		}
	case "exact":
		exEx := &ExactExpense{}
		if err = json.Unmarshal(data, exEx); err != nil {
			return err
		}
	case "percent":
		pEx := &PercentExpense{}
		if err = json.Unmarshal(data, pEx); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown expense type: %s", vs)
	}
	return nil
}

type ExpenseType string

const (
	Equal   ExpenseType = "equal"
	Exact   ExpenseType = "exact"
	Percent ExpenseType = "percent"
)
