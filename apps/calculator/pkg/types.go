package pkg

import "time"

type CalculationRecord struct {
	Name        string     `json:"name,omitempty"`
	Expression  string     `json:"expression" json:"expression,omitempty"`
	Result      float64    `json:"result" json:"result,omitempty"`
	SubmittedAt *time.Time `json:"submitted_at" json:"submitted_at,omitempty"`
}
