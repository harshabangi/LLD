package db

import "github.com/harshabangi/LLD/apps/splitwise/pkg"

type db struct {
	userMap      map[string]pkg.User
	balanceSheet map[string]map[string]int64
}

type DB interface {
	AddUser(u *pkg.User)
	AddExpense(e pkg.Expenser)
	ListBalances() map[string]int64
	GetBalance(userEmail string) int64
}

func NewDB() DB {
	return &db{
		userMap:      make(map[string]pkg.User),
		balanceSheet: make(map[string]map[string]int64),
	}
}

func (d *db) AddUser(u *pkg.User) {
	d.userMap[u.Email] = *u
	d.balanceSheet[u.Email] = map[string]int64{}
}

func (d *db) AddExpense(e pkg.Expenser) {
	paidBy := e.GetPaidBy()

	for _, v := range e.GetSplits() {
		if v.UserEmail == paidBy {
			continue
		}

		m := d.balanceSheet[paidBy]
		if _, ok := m[v.UserEmail]; !ok {
			m[v.UserEmail] = 0
		}
		m[v.UserEmail] += v.Amount
		d.balanceSheet[paidBy] = m

		m = d.balanceSheet[v.UserEmail]
		if _, ok := m[paidBy]; !ok {
			m[paidBy] = 0
		}
		m[v.UserEmail] -= v.Amount
		d.balanceSheet[v.UserEmail] = m
	}
}

func (d *db) ListBalances() map[string]int64 {
	balanceSheet := make(map[string]int64)

	for k1, v1 := range d.balanceSheet {
		var a int64 = 0
		for _, v2 := range v1 {
			a += v2
		}
		balanceSheet[k1] = a
	}
	return balanceSheet
}

func (d *db) GetBalance(userEmail string) int64 {
	var total int64 = 0
	for _, v := range d.balanceSheet[userEmail] {
		total += v
	}
	return total
}
