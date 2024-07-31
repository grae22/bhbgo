package accounts

type Account struct {
	Name        string
	AccountType string
	Parent      *Account
	Children    []*Account
	Balance     int
}
