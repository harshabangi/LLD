package memory

import (
	"fmt"
	"github.com/harshabangi/LLD/apps/calculator/pkg"
	"time"
)

type CalculationRecord struct {
	Name        string
	Expression  string
	Result      float64
	SubmittedAt *time.Time
}

func (c *CalculationRecord) ToModel() pkg.CalculationRecord {
	return pkg.CalculationRecord{
		Expression:  c.Expression,
		Result:      c.Result,
		SubmittedAt: c.SubmittedAt,
	}
}

type memoryStore struct {
	calculationsHistory []CalculationRecord
	calculationsMap     map[string]CalculationRecord
}

func NewMemoryStore() Store {
	return &memoryStore{
		calculationsHistory: []CalculationRecord{},
		calculationsMap:     make(map[string]CalculationRecord),
	}
}

type Store interface {
	ListCalculationRecords() []CalculationRecord
	AddCalculation(record CalculationRecord)
	GetCalculationRecord(recordName string) CalculationRecord
}

func (m *memoryStore) ListCalculationRecords() []CalculationRecord {
	return m.calculationsHistory
}

func (m *memoryStore) AddCalculation(record CalculationRecord) {
	m.calculationsHistory = append(m.calculationsHistory, record)

	if record.Name != "" {
		if _, ok := m.calculationsMap[record.Name]; ok {
			// for now, we will just panic, instead of returning error
			panic(fmt.Errorf("record with name '%s' already exists", record.Name))
		} else {
			m.calculationsMap[record.Name] = record
		}
	}
}

func (m *memoryStore) GetCalculationRecord(recordName string) CalculationRecord {
	return m.calculationsMap[recordName]
}
