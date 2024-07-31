package accounts

type rootAccounts struct {
	Bank      Account
	Income    Account
	Expense   Account
	Debtors   Account
	Creditors Account
}

var Roots rootAccounts
var RootsArr [5]*Account

func initialiseRegistry() {
	Roots.Bank = Account{
		Name:        "Bank",
		AccountType: AccountTypes[AccountTypeIndexBank],
		Children:    make([]*Account, 0),
	}

	Roots.Income = Account{
		Name:        "Income",
		AccountType: AccountTypes[AccountTypeIndexIncome],
		Children:    make([]*Account, 0),
	}

	Roots.Expense = Account{
		Name:        "Expense",
		AccountType: AccountTypes[AccountTypeIndexExpense],
		Children:    make([]*Account, 0),
	}

	Roots.Debtors = Account{
		Name:        "Debtors",
		AccountType: AccountTypes[AccountTypeIndexDebtor],
		Children:    make([]*Account, 0),
	}

	Roots.Creditors = Account{
		Name:        "Creditors",
		AccountType: AccountTypes[AccountTypeIndexCreditor],
		Children:    make([]*Account, 0),
	}

	RootsArr[AccountTypeIndexBank] = &Roots.Bank
	RootsArr[AccountTypeIndexIncome] = &Roots.Income
	RootsArr[AccountTypeIndexExpense] = &Roots.Expense
	RootsArr[AccountTypeIndexCreditor] = &Roots.Debtors
	RootsArr[AccountTypeIndexDebtor] = &Roots.Creditors
}
