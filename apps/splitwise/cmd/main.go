package main

import (
	"fmt"
	"github.com/harshabangi/LLD/apps/splitwise/internal/db"
	"github.com/harshabangi/LLD/apps/splitwise/internal/service"
	"github.com/harshabangi/LLD/apps/splitwise/pkg"
	"log"
)

func main() {
	h := service.NewHandler(db.NewDB())

	h.AddUser(&pkg.User{Email: "u1"})
	h.AddUser(&pkg.User{Email: "u2"})
	h.AddUser(&pkg.User{Email: "u3"})

	err := h.AddExpense(&pkg.EqualExpense{
		Expense: pkg.Expense{Type: "equal", Description: "desc1", Amount: 100, PaidBy: "u2"},
		Users:   []string{"u2", "u3"},
	})
	if err != nil {
		log.Fatal(err)
	}

	err = h.AddExpense(&pkg.ExactExpense{
		Expense: pkg.Expense{Type: "exact", Description: "desc2", Amount: 100, PaidBy: "u1"},
		ExactSplits: []pkg.Split{
			{UserEmail: "u1", Amount: 30},
			{UserEmail: "u2", Amount: 70},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(h.ListBalances())

}
