package BankAccount

type Account struct {
	Holder  string
	Balance float64
}

func CreateAccout(name string, deposit float64) Account {
	return Account{
		Holder:  name,
		Balance: deposit,
	}
}
func (a *Account) Deposit(deposit float64) {
	a.Balance = a.Balance + deposit
}
func (a *Account) Withdraw(withdraw float64) {
	a.Balance = a.Balance - withdraw
}
