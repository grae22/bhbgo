package accounts

type Account struct {
	Name        string
	AccountType string
	Children    []*Account
	Balance     Balance
}

type Balance struct {
	ByPeriod []int
}
