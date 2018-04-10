package account

type Account struct {
	balance int64
	closed  bool
	actionc chan func()
	quitc   chan chan struct{}
}

func Open(amt int64) *Account {
	if amt < 0 {
		return nil
	}

	a := &Account{
		balance: amt,
		closed:  false,
		actionc: make(chan func()),
		quitc:   make(chan chan struct{}),
	}
	go a.loop()
	return a
}

func (a *Account) loop() {
	for {
		select {
		case f := <-a.actionc:
			f()
		case q := <-a.quitc:
			close(q)
			return
		}
	}
}

func (a *Account) Balance() (balance int64, ok bool) {
	b := make(chan int64)
	okc := make(chan bool)

	a.actionc <- func() {
		if a.closed {
			b <- 0
			okc <- false
		} else {
			b <- a.balance
			okc <- true
		}
	}
	return <-b, <-okc

}

func (a *Account) Close() (payout int64, ok bool) {
	p := make(chan int64)
	okc := make(chan bool)

	a.actionc <- func() {
		if a.closed {
			p <- 0
			okc <- false
		} else {
			p <- a.balance
			// Update the balance
			a.balance = 0
			a.closed = true
			okc <- true
		}
	}
	return <-p, <-okc
}

func (a *Account) Deposit(amt int64) (newbalance int64, ok bool) {
	nb := make(chan int64)
	okc := make(chan bool)

	a.actionc <- func() {
		if a.balance+amt < 0 || a.closed {
			nb <- 0
			okc <- false
		} else {
			a.balance += amt

			nb <- a.balance
			okc <- true
		}
	}
	return <-nb, <-okc
}
