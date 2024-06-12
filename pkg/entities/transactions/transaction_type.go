package entities

type TransactionType int

const (
    TransactionTypeDeposit TransactionType = iota + 1
    TransactionTypeWithdrawal
    TransactionTypeTransfer
    TransactionTypePayment
    TransactionTypePurchase
    TransactionTypeSalary
    TransactionTypeExpense
    TransactionTypeLoan
    TransactionTypeInterest
)