package account

import "sync"

type Account struct {
	balance int64
	closed  bool
	mu      *sync.Mutex
}

func Open(amt int64) *Account {
	if amt < 0 {
		return nil
	}
	return &Account{
		balance: amt,
		closed:  false,
		mu:      &sync.Mutex{},
	}
}

func (a *Account) Balance() (balance int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.closed {
		return 0, false
	}
	return a.balance, true
}
func (a *Account) Close() (payout int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.closed {
		return 0, false
	}

	balance := a.balance

	// Update the account details.
	a.balance = 0
	a.closed = true

	return balance, true
}
func (a *Account) Deposit(amt int64) (newbalance int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.balance < -amt {
		return 0, false
	}

	if a.closed {
		return 0, false
	}
	a.balance += amt
	return a.balance, true
}
