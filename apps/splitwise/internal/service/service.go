package service

import (
	"github.com/harshabangi/LLD/apps/splitwise/internal/db"
	"github.com/harshabangi/LLD/apps/splitwise/pkg"
)

type Handler struct {
	db db.DB
}

func NewHandler(db db.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) AddUser(u *pkg.User) {
	h.db.AddUser(u)
}

func (h *Handler) AddExpense(e pkg.Expenser) error {
	if err := e.Validate(); err != nil {
		return err
	}
	h.db.AddExpense(e)
	return nil
}

func (h *Handler) ListBalances() map[string]int64 {
	return h.db.ListBalances()
}

func (h *Handler) GetUserBalance(userEmail string) int64 {
	return h.db.GetBalance(userEmail)
}
